package markdown

import (
	"bytes"
	"github.com/sascha-andres/people2md/internal/types"
	"text/template"
)

func (mdData *MarkdownData) BuildAddresses(c *types.Contact, addresses *template.Template) string {
	if len(c.Addresses) > 0 {
		var buff bytes.Buffer
		if err := addresses.Execute(&buff, c.Addresses); err != nil {
			return err.Error()
		} else {
			return buff.String()
		}
	}
	return ""
}
