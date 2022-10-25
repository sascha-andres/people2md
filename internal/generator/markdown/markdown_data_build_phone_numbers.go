package markdown

import (
	"bytes"
	"github.com/sascha-andres/people2md/internal/types"
	"text/template"
)

func (mdData *MarkdownData) BuildPhoneNumbers(c *types.Contact, phoneNumbers *template.Template) string {
	if len(c.PhoneNumbers) > 0 {
		var buff bytes.Buffer
		if err := phoneNumbers.Execute(&buff, c.PhoneNumbers); err != nil {
			return err.Error()
		} else {
			return buff.String()
		}
	}
	return ""
}
