package internal

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

func (mdData *MarkdownData) WriteMarkdown(outer *template.Template, c *Contact) {
	var buff bytes.Buffer
	outer.Execute(&buff, mdData)
	var fileName = ""
	if len(c.Names) > 0 {
		fileName = toFileName(c.Names[0].DisplayName) + ".md"
	} else {
		fileName = toFileName(c.Organizations[0].Name) + ".md"
	}

	ioutil.WriteFile(fileName, buff.Bytes(), 0600)
}

