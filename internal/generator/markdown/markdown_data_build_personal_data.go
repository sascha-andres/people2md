package markdown

import (
	"bytes"
	"github.com/sascha-andres/people2md/internal/types"
	"text/template"
)

func (mdData *MarkdownData) BuildPersonalData(personalData *template.Template, c *types.Contact) string {
	var pd bytes.Buffer
	if err := personalData.Execute(&pd, c); err != nil {
		return err.Error()
	} else {
		return pd.String()
	}
}
