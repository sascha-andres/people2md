package internal

import (
	"bytes"
	"text/template"
)

func (mdData *MarkdownData) BuildPersonalData(personalData *template.Template, c *Contact) {
	var pd bytes.Buffer
	if err := personalData.Execute(&pd, c); err != nil {
		mdData.PersonalData = err.Error()
	} else {
		mdData.PersonalData = pd.String()
	}
}
