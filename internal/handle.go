package internal

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"slices"

	"github.com/sascha-andres/people2md/internal/types"
	"github.com/sascha-andres/sbrdata/v2"
)

// sanitizePhoneNumber removes all characters that should not be part of a phone number
// and tries to format the number to a German number
func sanitizePhoneNumber(number string) string {
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")
	number = strings.ReplaceAll(number, "/", "")
	number = strings.ReplaceAll(number, "+49", "0")
	number = strings.ReplaceAll(number, "+", "00")
	number = strings.ReplaceAll(number, "(", "")
	number = strings.ReplaceAll(number, ")", "")
	return strings.TrimSpace(number)
}

// filterCalls returns a list of calls that are related to the contact
func filterCalls(c types.Contact, allCalls []sbrdata.Call) []sbrdata.Call {
	var result = make([]sbrdata.Call, 0)

	for _, call := range allCalls {
		include := false
		for i := range c.Names {
			include = c.Names[i].DisplayName == call.GetContactName()
			if include {
				break
			}
		}
		if !include {
			for i := range c.Organizations {
				include = c.Organizations[i].Name == call.GetContactName()
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
			for i := range c.PhoneNumbers {
				compare := sanitizePhoneNumber(c.PhoneNumbers[i].Value)
				//if strings.Contains(c.PhoneNumbers[i].Value, "31454") {
				//	log.Printf("compare: %s, num: %s", compare, num)
				//}
				include = strings.HasSuffix(compare, num)
				if include {
					break
				}
			}
		}
		if include {
			result = append(result, call)
		}
	}
	return result
}

// handle is running the generation process
func (app *Application) handle(data types.DataReferences, generator types.DataBuilder, templates *types.Templates) {
	e := &types.Elements{}

	e.ETag = data.Contact.Etag
	e.ResourceName = data.Contact.ResourceName
	if len(data.Contact.Birthdays) > 0 {
		e.Birthday = types.TemplateDate(data.Contact.Birthdays[0].Date.Year, data.Contact.Birthdays[0].Date.Month, data.Contact.Birthdays[0].Date.Day)
	}

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
	calls, err := data.Collection.AllCalls()
	if err != nil {
		// TODO log...
		return
	}
	filteredCalls := filterCalls(*data.Contact, calls)
	slices.SortFunc(filteredCalls, func(i, j sbrdata.Call) int {
		a, err := strconv.Atoi(i.GetDate())
		if err != nil {
			return 0
		}
		b, err := strconv.Atoi(j.GetDate())
		if err != nil {
			return 0
		}
		return b - a
	})
	e.Calls = filteredCalls
	sms, err := data.Collection.AllSms()
	if err != nil {
		// TODO log...
		return
	}
	ml = addSmsToList(data.Contact, &sbrdata.Messages{Sms: sms}, ml)
	mms, err := data.Collection.AllMms()
	if err != nil {
		// TODO log...
		return
	}
	ml = addMmsToList(data.Contact, &sbrdata.Messages{Mms: mms}, ml)
	slices.SortFunc(ml, func(i, j types.Message) int {
		if i.UnixTimestamp < j.UnixTimestamp {
			return -1
		}
		if i.UnixTimestamp > j.UnixTimestamp {
			return 1
		}
		return 0
	})
	e.MessageData = ml
	e.PersonalData = generator.BuildPersonalData(templates.PersonalData, data.Contact)
	e.Tags = generator.BuildTags(data.Tags, data.TagPrefix, data.Contact, data.Groups)
	e.Addresses = generator.BuildAddresses(data.Contact, templates.Addresses)
	e.PhoneNumbers = generator.BuildPhoneNumbers(data.Contact, templates.PhoneNumbers)
	e.Email = generator.BuildEmailAddresses(data.Contact, templates.EmailAddresses)
	additionalDataPath := path.Join(data.PathForFiles, e.MainLinkName)
	if _, err := os.Stat(additionalDataPath); os.IsNotExist(err) {
		_ = os.MkdirAll(additionalDataPath, 0770)
	}
	if len(e.Calls) > 0 && templates.Calls != nil {
		var (
			buff bytes.Buffer
		)
		_ = templates.Calls.Execute(&buff, e)
		app.writeBufferToFile(buff, path.Join(additionalDataPath, e.MainLinkName+" Calls"))
	}
	if len(e.MessageData) > 0 && templates.Messages != nil {
		var (
			buff bytes.Buffer
		)
		_ = templates.Messages.Execute(&buff, e)
		app.writeBufferToFile(buff, path.Join(additionalDataPath, e.MainLinkName+" Messages"))
	}

	var mainContactSheetBuffer bytes.Buffer
	_ = templates.ContactSheet.Execute(&mainContactSheetBuffer, e)
	app.writeBufferToFile(mainContactSheetBuffer, path.Join(data.PathForFiles, e.MainLinkName))

	destinationPath := path.Join(additionalDataPath, e.MainLinkName+" Notes")
	if _, err := os.Stat(destinationPath + "." + app.fileExtension); errors.Is(err, os.ErrNotExist) {
		var notesSheetBuffer bytes.Buffer
		_ = templates.NotesSheet.Execute(&notesSheetBuffer, e)
		app.writeBufferToFile(notesSheetBuffer, destinationPath)
	}
}

// writeBufferToFile writes the buffer to a file
func (app *Application) writeBufferToFile(buff bytes.Buffer, destinationPath string) {
	var res int = 1
	hasher := sha256.New()
	hasher.Write(buff.Bytes())
	hashNew := hasher.Sum(nil)

	destinationPathWithExtension := destinationPath + "." + app.fileExtension
	fileData, err := os.ReadFile(destinationPathWithExtension)
	if err == nil {
		hasher.Reset()

		hasher.Write(fileData)
		hashOld := hasher.Sum(nil)

		res = bytes.Compare(hashOld, hashNew)
	} else {
		log.Printf("could not read existing file: %s", err)
	}

	if res == 0 {
		if app.verbose {
			log.Printf("identical: %s", destinationPathWithExtension)
		}
		return
	}

	if app.dryRun {
		log.Printf("would write %s", destinationPathWithExtension)
		return
	}

	log.Printf("replacing %s", destinationPathWithExtension)

	err = os.WriteFile(destinationPathWithExtension, buff.Bytes(), 0600)
	if err != nil {
		log.Printf("could not write file: %s", err)
		return
	}
}

// addMmsToList adds mms messages to the message list
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

// sanitizeBody sanitizes the body of a message
func sanitizeBody(body string) string {
	result := strings.Replace(body, "<", "&lt;", -1)
	result = strings.Replace(result, ">", "&gt;", -1)
	result = strings.Replace(result, "\n", "<br />", -1)
	result = strings.Replace(result, "[", "", -1)
	if strings.HasPrefix(result, "0") {
		result = result[1:]
	}
	return result
}

// addSmsToList adds sms messages to the message list
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
