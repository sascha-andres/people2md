package internal

import (
	"bytes"
	"text/template"
)

func (mdData *MarkdownData) BuildAddresses(c *Contact, addresses *template.Template) {
	if len(c.Addresses) > 0 {
		var buff bytes.Buffer
		if err := addresses.Execute(&buff, c.Addresses); err != nil {
			mdData.Addresses = err.Error()
		} else {
			mdData.Addresses = buff.String()
		}
	}
}

