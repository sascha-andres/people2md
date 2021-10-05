package internal

import (
	"bytes"
	"text/template"
)

func (mdData *MarkdownData) BuildPhoneNumbers(c *Contact, phoneNumbers *template.Template) {
	if len(c.PhoneNumbers) > 0 {
		var buff bytes.Buffer
		if err := phoneNumbers.Execute(&buff, c.PhoneNumbers); err != nil {
			mdData.PhoneNumbers = err.Error()
		} else {
			mdData.PhoneNumbers = buff.String()
		}
	}
}
