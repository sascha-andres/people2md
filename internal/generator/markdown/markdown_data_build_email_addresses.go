package markdown

import (
	"bytes"
	"github.com/sascha-andres/people2md/internal/types"
	"text/template"
)

func (mdData *MarkdownData) BuildEmailAddresses(c *types.Contact, emailAddresses *template.Template) string {
	if len(c.EmailAddresses) > 0 {
		var buff bytes.Buffer
		if err := emailAddresses.Execute(&buff, c.EmailAddresses); err != nil {
			return err.Error()
		} else {
			return buff.String()
		}
	}
	return ""
}
