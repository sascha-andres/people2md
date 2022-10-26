package internal

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/sascha-andres/people2md/internal/types"
	"github.com/sascha-andres/sbrdata"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"text/template"
)

func Handle(c *types.Contact, generator types.DataBuilder, pathForFiles, tags string, personalData *template.Template, groups []types.ContactGroup, addresses, phoneNumbers, emailAddresses, outer *template.Template, sms sbrdata.Messages, calls sbrdata.Calls, verbose bool) {
	e := &types.Elements{}

	e.ETag = c.Etag
	e.ResourceName = c.ResourceName

	generator.SetETag(c.Etag)
	generator.SetResourceName(c.ResourceName)

	// build message list
	var ml types.MessageList
	ml = addSmsToList(c, sms, ml)
	ml = addMmsToList(c, sms, ml)
	sort.Sort(ml)

	e.Calls = generator.BuildCalls(calls, c)
	e.Sms = generator.BuildSms(ml)
	e.PersonalData = generator.BuildPersonalData(personalData, c)
	e.Tags = generator.BuildTags(tags, c, groups)
	e.Addresses = generator.BuildAddresses(c, addresses)
	e.PhoneNumbers = generator.BuildPhoneNumbers(c, phoneNumbers)
	e.Email = generator.BuildEmailAddresses(c, emailAddresses)

	var buff bytes.Buffer
	outer.Execute(&buff, e)
	var fileName = ""
	if len(c.Names) > 0 {
		fileName = toFileName(c.Names[0].DisplayName) + ".md"
	} else {
		fileName = toFileName(c.Organizations[0].Name) + ".md"
	}

	destinationPath := path.Join(pathForFiles, fileName)

	if _, err := os.Stat(destinationPath); errors.Is(err, os.ErrNotExist) {
		os.WriteFile(destinationPath, buff.Bytes(), 0600)
		return
	}

	hasher := sha256.New()
	hasher.Write(buff.Bytes())
	hashNew := hasher.Sum(nil)

	fileData, err := os.ReadFile(destinationPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read existing file: %s", err)
		return
	}

	hasher = sha256.New()
	hasher.Write(fileData)
	hashOld := hasher.Sum(nil)

	res := bytes.Compare(hashOld, hashNew)

	if res == 0 {
		return
	}

	if verbose {
		log.Printf("replacing %s", destinationPath)
	}
	err = os.WriteFile(destinationPath, buff.Bytes(), 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not write file: %s", err)
		return
	}
}

func addMmsToList(c *types.Contact, sms sbrdata.Messages, ml types.MessageList) types.MessageList {
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

func addSmsToList(c *types.Contact, sms sbrdata.Messages, ml types.MessageList) types.MessageList {
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
