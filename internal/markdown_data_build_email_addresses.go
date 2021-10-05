package internal

import (
	"bytes"
	"text/template"
)

func (mdData *MarkdownData) BuildEmailAddresses(c *Contact, emailAddresses *template.Template) {
	if len(c.EmailAddresses) > 0 {
		var buff bytes.Buffer
		if err := emailAddresses.Execute(&buff, c.EmailAddresses); err != nil {
			mdData.Email = err.Error()
		} else {
			mdData.Email = buff.String()
		}
	}
}
