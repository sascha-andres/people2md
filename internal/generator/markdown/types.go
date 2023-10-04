package markdown

import (
	"embed"
	"text/template"

	"github.com/sascha-andres/people2md/internal/types"
)

//go:embed content/*
var content embed.FS

type (
	// MarkdownData is the data structure for the markdown generator (Frontmatter, etc.)
	MarkdownData struct {
		ETag         string
		ResourceName string
	}
)

func (mdData *MarkdownData) SetETag(etag string) {
	mdData.ETag = etag
}

func (mdData *MarkdownData) SetResourceName(rn string) {
	mdData.ResourceName = rn
}

func (mdData *MarkdownData) GetTemplate(id types.TemplateIdentifier) *template.Template {
	funcMap := template.FuncMap{
		"replace": types.TemplateReplace,
	}
	templateContent := ""

	switch id {
	case types.ContactSheetTemplate:
		templateContent = loadFromFSAsString("content/contactsheet.tmpl")
	case types.AddressesTemplate:
		templateContent = loadFromFSAsString("content/addresses.tmpl")
	case types.PersonalDataTemplate:
		templateContent = loadFromFSAsString("content/personal.tmpl")
	case types.PhoneNumbersTemplate:
		templateContent = loadFromFSAsString("content/phone.tmpl")
	case types.EmailAddressesTemplate:
		templateContent = loadFromFSAsString("content/emails.tmpl")
	case types.CallsTemplate:
		templateContent = loadFromFSAsString("content/calls.tmpl")
	case types.MessagesTemplate:
		templateContent = loadFromFSAsString("content/messages.tmpl")
	case types.NotesSheetTemplate:
		templateContent = loadFromFSAsString("content/notes.tmpl")
	}

	if templateContent == "" {
		return nil
	}

	return template.Must(template.New("outer").Funcs(funcMap).Parse(templateContent))
}

func loadFromFSAsString(filename string) string {
	return string(loadFromFS(filename))
}

func loadFromFS(filename string) []byte {
	data, err := content.ReadFile(filename)
	if err != nil {
		return nil
	}
	return data
}

func (mdData *MarkdownData) GetTemplateData(id types.TemplateIdentifier) []byte {
	switch id {
	case types.ContactSheetTemplate:
		return loadFromFS("content/contactsheet.tmpl")
	case types.AddressesTemplate:
		return loadFromFS("content/addresses.tmpl")
	case types.PersonalDataTemplate:
		return loadFromFS("content/personal.tmpl")
	case types.PhoneNumbersTemplate:
		return loadFromFS("content/phone.tmpl")
	case types.EmailAddressesTemplate:
		return loadFromFS("content/emails.tmpl")
	case types.CallsTemplate:
		return loadFromFS("content/calls.tmpl")
	case types.MessagesTemplate:
		return loadFromFS("content/messages.tmpl")
	case types.NotesSheetTemplate:
		return loadFromFS("content/notes.tmpl")
	}

	return nil
}
