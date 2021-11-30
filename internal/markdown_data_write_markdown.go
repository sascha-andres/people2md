package internal

import (
	"bytes"
	"io/ioutil"
	"path"
	"text/template"
)

func (mdData *MarkdownData) WriteMarkdown(pathForFiles string, outer *template.Template, c *Contact) {
	var buff bytes.Buffer
	outer.Execute(&buff, mdData)
	var fileName = ""
	if len(c.Names) > 0 {
		fileName = toFileName(c.Names[0].DisplayName) + ".md"
	} else {
		fileName = toFileName(c.Organizations[0].Name) + ".md"
	}

	ioutil.WriteFile(path.Join(pathForFiles, fileName), buff.Bytes(), 0600)
}
