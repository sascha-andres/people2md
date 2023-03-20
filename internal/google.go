package internal

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"text/template"

	"github.com/sascha-andres/people2md/internal/types"
	"github.com/sascha-andres/sbrdata"
)

func Handle(c *types.Contact, generator types.DataBuilder, pathForFiles, tags, tagPrefix string, personalData *template.Template, groups []types.ContactGroup, addresses, phoneNumbers, emailAddresses, outer *template.Template, sms *sbrdata.Messages, callData *sbrdata.Calls, collection *sbrdata.Collection, verbose bool, calls, messages *template.Template) {
	e := &types.Elements{}

	e.ETag = c.Etag
	e.ResourceName = c.ResourceName

	generator.SetETag(c.Etag)
	generator.SetResourceName(c.ResourceName)

	// build message list
	var ml types.MessageList
	if sms != nil {
		ml = addSmsToList(c, sms, ml)
		ml = addMmsToList(c, sms, ml)
	}

	if len(c.Names) > 0 {
		e.MainLinkName = toFileName(c.Names[0].DisplayName)
	} else {
		e.MainLinkName = toFileName(c.Organizations[0].Name)
	}
	if callData != nil {
		e.Calls = generator.BuildCalls(callData, c)
	}
	if collection != nil {
		e.Calls = generator.BuildCalls(&sbrdata.Calls{
			Call: collection.Calls,
		}, c)
		ml = addSmsToList(c, &sbrdata.Messages{Sms: collection.SMS}, ml)
		ml = addMmsToList(c, &sbrdata.Messages{Mms: collection.MMS}, ml)
	}
	if len(ml) > 0 {
		sort.Sort(ml)
		e.Messages = generator.BuildMessages(ml)
	}
	e.PersonalData = generator.BuildPersonalData(personalData, c)
	e.Tags = generator.BuildTags(tags, tagPrefix, c, groups)
	e.Addresses = generator.BuildAddresses(c, addresses)
	e.PhoneNumbers = generator.BuildPhoneNumbers(c, phoneNumbers)
	e.Email = generator.BuildEmailAddresses(c, emailAddresses)
	if len(e.Calls) > 0 {
		var (
			buff bytes.Buffer
		)
		_ = calls.Execute(&buff, e)
		writeBufferToFile(path.Join(pathForFiles, e.MainLinkName+" Calls.md"), buff, verbose)
	}
	if len(e.Messages) > 0 {
		var (
			buff bytes.Buffer
		)
		_ = messages.Execute(&buff, e)
		writeBufferToFile(path.Join(pathForFiles, e.MainLinkName+" Messages.md"), buff, verbose)
	}

	var buff bytes.Buffer
	_ = outer.Execute(&buff, e)
	writeBufferToFile(path.Join(pathForFiles, e.MainLinkName+".md"), buff, verbose)
}

func writeBufferToFile(destinationPath string, buff bytes.Buffer, verbose bool) {
	if _, err := os.Stat(destinationPath); errors.Is(err, os.ErrNotExist) {
		_ = os.WriteFile(destinationPath, buff.Bytes(), 0600)
		return
	}

	var res int = 1
	hasher := sha256.New()
	hasher.Write(buff.Bytes())
	hashNew := hasher.Sum(nil)

	fileData, err := os.ReadFile(destinationPath)
	if err == nil {
		hasher.Reset()

		hasher.Write(fileData)
		hashOld := hasher.Sum(nil)

		res = bytes.Compare(hashOld, hashNew)
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "could not read existing file: %s", err)
	}

	if res == 0 {
		if verbose {
			log.Printf("identical: %s", destinationPath)
		}
		return
	}

	if verbose {
		log.Printf("replacing %s", destinationPath)
	}
	err = os.WriteFile(destinationPath, buff.Bytes(), 0600)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "could not write file: %s", err)
		return
	}
}

func addMmsToList(c *types.Contact, sms *sbrdata.Messages, ml types.MessageList) types.MessageList {
	for _, message := range sms.Mms {
		include := false
		for _, name := range c.Names {
			include = name.DisplayName == message.ContactName
			if include {
				break
			}
		}
		if !include {
			for _, org := range c.Organizations {
				include = org.Name == message.ContactName
				if include {
					break
				}
			}
		}
		if include {
			d, err := strconv.ParseUint(message.Date, 10, 64)
			if err != nil {
				continue
			}
			direction := "received"
			switch message.MsgBox {
			case "1":
				break
			case "2":
				direction = "sent"
				break
			default:
				continue
			}
			body := "not - found"
			for _, part := range message.Parts.Part {
				if part.Ct == "text/plain" || part.Name == "body" {
					body = part.AttrText
					break
				}
			}
			ml = append(ml, types.Message{
				UnixTimestamp: d,
				Date:          message.ReadableDate,
				Direction:     direction,
				Text:          body,
			})
		}
	}
	return ml
}

func addSmsToList(c *types.Contact, sms *sbrdata.Messages, ml types.MessageList) types.MessageList {
	for _, message := range sms.Sms {
		include := false
		for _, name := range c.Names {
			include = name.DisplayName == message.ContactName
			if include {
				break
			}
		}
		if !include {
			for _, org := range c.Organizations {
				include = org.Name == message.ContactName
				if include {
					break
				}
			}
		}
		if include {
			d, err := strconv.ParseUint(message.Date, 10, 64)
			if err != nil {
				continue
			}
			direction := "received"
			if message.Type == "2" {
				direction = "sent"
			}
			ml = append(ml, types.Message{
				UnixTimestamp: d,
				Date:          message.ReadableDate,
				Direction:     direction,
				Text:          message.Body,
			})
		}
	}
	return ml
}
