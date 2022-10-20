package internal

import (
	"fmt"
	"github.com/sascha-andres/sbrdata"
	"strings"
)

func (mdData *MarkdownData) BuildSms(sms sbrdata.Messages, c *Contact) {
	result := ""

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
			if result == "" {
				result = `|Date|Direction|Texta|
|---|---|---|`
			}
			direction := "received"
			if message.Type == "2" {
				direction = "sent"
			}
			result = fmt.Sprintf("%s\n|%s|%s|%s|", result, message.ReadableDate, direction, sanitizeBody(message.Body))
		}
	}

	mdData.Sms = result
}

func sanitizeBody(body string) string {
	result := strings.Replace(body, "<", "&lt;", -1)
	result = strings.Replace(result, ">", "&gt;", -1)
	result = strings.Replace(result, "\n", "<br />", -1)
	return result
}
