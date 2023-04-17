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
	"strings"

	"github.com/sascha-andres/people2md/internal/types"
	"github.com/sascha-andres/sbrdata"
)

func (app *Application) handle(data types.DataReferences, generator types.DataBuilder, templates *types.Templates) {
	e := &types.Elements{}

	e.ETag = data.Contact.Etag
	e.ResourceName = data.Contact.ResourceName

	if 0 == len(data.Contact.Names) && 0 == len(data.Contact.Organizations) {
		return
	}
	if len(data.Contact.Names) > 0 {
		data.Contact.Names[0].DisplayName = strings.TrimSpace(data.Contact.Names[0].DisplayName)
	}
	if len(data.Contact.Organizations) > 0 {
		data.Contact.Organizations[0].Name = strings.TrimSpace(data.Contact.Organizations[0].Name)
	}

	generator.SetETag(data.Contact.Etag)
	generator.SetResourceName(data.Contact.ResourceName)

	var ml types.MessageList

	if len(data.Contact.Names) > 0 {
		e.MainLinkName = toFileName(data.Contact.Names[0].DisplayName)
	} else {
		e.MainLinkName = toFileName(data.Contact.Organizations[0].Name)
	}
	if data.Collection != nil {
		e.Calls = generator.BuildCalls(&sbrdata.Calls{
			Call: data.Collection.Calls,
		}, data.Contact)
		ml = addSmsToList(data.Contact, &sbrdata.Messages{Sms: data.Collection.SMS}, ml)
		ml = addMmsToList(data.Contact, &sbrdata.Messages{Mms: data.Collection.MMS}, ml)
	} else {
		if data.CallData != nil {
			e.Calls = generator.BuildCalls(data.CallData, data.Contact)
		}
		if data.Sms != nil {
			ml = addSmsToList(data.Contact, data.Sms, ml)
			ml = addMmsToList(data.Contact, data.Sms, ml)
		}
	}
	if len(ml) > 0 {
		sort.Sort(ml)
		e.Messages = generator.BuildMessages(ml)
	}
	e.PersonalData = generator.BuildPersonalData(templates.PersonalData, data.Contact)
	e.Tags = generator.BuildTags(data.Tags, data.TagPrefix, data.Contact, data.Groups)
	e.Addresses = generator.BuildAddresses(data.Contact, templates.Addresses)
	e.PhoneNumbers = generator.BuildPhoneNumbers(data.Contact, templates.PhoneNumbers)
	e.Email = generator.BuildEmailAddresses(data.Contact, templates.EmailAddresses)
	if len(e.Calls) > 0 {
		var (
			buff bytes.Buffer
		)
		_ = templates.Calls.Execute(&buff, e)
		app.writeBufferToFile(path.Join(data.PathForFiles, e.MainLinkName+" Calls.md"), buff)
	}
	if len(e.Messages) > 0 {
		var (
			buff bytes.Buffer
		)
		_ = templates.Messages.Execute(&buff, e)
		app.writeBufferToFile(path.Join(data.PathForFiles, e.MainLinkName+" Messages.md"), buff)
	}

	var buff bytes.Buffer
	_ = templates.Outer.Execute(&buff, e)
	app.writeBufferToFile(path.Join(data.PathForFiles, e.MainLinkName+".md"), buff)
}

func (app *Application) writeBufferToFile(destinationPath string, buff bytes.Buffer) {
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
		if app.verbose {
			log.Printf("identical: %s", destinationPath)
		}
		return
	}

	log.Printf("replacing %s", destinationPath)

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
