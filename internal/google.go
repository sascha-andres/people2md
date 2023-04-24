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

func sanitizePhoneNumber(number string) string {
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")
	number = strings.ReplaceAll(number, "/", "")
	number = strings.ReplaceAll(number, "+49", "0")
	number = strings.ReplaceAll(number, "+", "00")
	number = strings.ReplaceAll(number, "(0)", "")
	return strings.TrimSpace(number)
}

func filterCalls(c types.Contact, allCalls *sbrdata.Calls) *sbrdata.Calls {
	var result = &sbrdata.Calls{}
	for _, call := range allCalls.Call {
		include := false
		for _, name := range c.Names {
			include = name.DisplayName == call.GetContactName()
			if include {
				break
			}
		}
		if !include {
			for _, org := range c.Organizations {
				include = org.Name == call.GetContactName()
				if include {
					break
				}
			}
		}
		// TODO: check whether it is feasible to
		//  include sth like contains with a cut off of
		//  the last number (having 1234560 as the central
		//  number but identify 12345678 also for this
		//  contact
		if call.GetNumber() != "" {
			num := sanitizePhoneNumber(call.GetNumber())
			if strings.HasPrefix(num, "0") {
				num = num[1:]
			}
			for _, p := range c.PhoneNumbers {
				compare := sanitizePhoneNumber(p.Value)
				include = strings.HasSuffix(compare, num)
				if include {
					break
				}
			}
		}
		if include {
			result.Call = append(result.Call, call)
		}
	}
	return result
}

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
		filteredCalls := filterCalls(*data.Contact, &sbrdata.Calls{
			Call: data.Collection.Calls,
		})
		e.CallData = filteredCalls
		ml = addSmsToList(data.Contact, &sbrdata.Messages{Sms: data.Collection.SMS}, ml)
		ml = addMmsToList(data.Contact, &sbrdata.Messages{Mms: data.Collection.MMS}, ml)
	} else {
		if data.CallData != nil {
			e.CallData = data.CallData
		}
		if data.Sms != nil {
			ml = addSmsToList(data.Contact, data.Sms, ml)
			ml = addMmsToList(data.Contact, data.Sms, ml)
		}
	}
	sort.Sort(ml)
	e.MessageData = ml
	e.PersonalData = generator.BuildPersonalData(templates.PersonalData, data.Contact)
	e.Tags = generator.BuildTags(data.Tags, data.TagPrefix, data.Contact, data.Groups)
	e.Addresses = generator.BuildAddresses(data.Contact, templates.Addresses)
	e.PhoneNumbers = generator.BuildPhoneNumbers(data.Contact, templates.PhoneNumbers)
	e.Email = generator.BuildEmailAddresses(data.Contact, templates.EmailAddresses)
	if len(e.CallData.Call) > 0 && templates.Calls != nil {
		var (
			buff bytes.Buffer
		)
		_ = templates.Calls.Execute(&buff, e)
		app.writeBufferToFile(path.Join(data.PathForFiles, e.MainLinkName+" Calls"), buff)
	}
	if len(e.MessageData) > 0 && templates.Messages != nil {
		var (
			buff bytes.Buffer
		)
		_ = templates.Messages.Execute(&buff, e)
		app.writeBufferToFile(path.Join(data.PathForFiles, e.MainLinkName+" Messages"), buff)
	}

	var buff bytes.Buffer
	_ = templates.ContactSheet.Execute(&buff, e)
	app.writeBufferToFile(path.Join(data.PathForFiles, e.MainLinkName), buff)
}

func (app *Application) writeBufferToFile(destinationPath string, buff bytes.Buffer) {
	destinationPathWithExtension := destinationPath + "." + app.fileExtension
	if _, err := os.Stat(destinationPathWithExtension); errors.Is(err, os.ErrNotExist) {
		_ = os.WriteFile(destinationPathWithExtension, buff.Bytes(), 0600)
		return
	}

	var res int = 1
	hasher := sha256.New()
	hasher.Write(buff.Bytes())
	hashNew := hasher.Sum(nil)

	fileData, err := os.ReadFile(destinationPathWithExtension)
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
			log.Printf("identical: %s", destinationPathWithExtension)
		}
		return
	}

	log.Printf("replacing %s", destinationPathWithExtension)

	err = os.WriteFile(destinationPathWithExtension, buff.Bytes(), 0600)
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
				Text:          sanitizeBody(body),
			})
		}
	}
	return ml
}

func sanitizeBody(body string) string {
	result := strings.Replace(body, "<", "&lt;", -1)
	result = strings.Replace(result, ">", "&gt;", -1)
	result = strings.Replace(result, "\n", "<br />", -1)
	result = strings.Replace(result, "[", "", -1)
	return result
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
				Text:          sanitizeBody(message.Body),
			})
		}
	}
	return ml
}
