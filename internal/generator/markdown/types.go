package markdown

import (
	"text/template"

	"github.com/sascha-andres/people2md/internal/types"
)

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
	switch id {
	case types.ContactSheetTemplate:
		return template.Must(template.New("outer").Funcs(funcMap).Parse(ContactSheetTemplate))
	case types.AddressesTemplate:
		return template.Must(template.New("outer").Funcs(funcMap).Parse(AddressesTemplate))
	case types.PersonalDataTemplate:
		return template.Must(template.New("outer").Funcs(funcMap).Parse(PersonalDataTemplate))
	case types.PhoneNumbersTemplate:
		return template.Must(template.New("outer").Funcs(funcMap).Parse(PhoneNumbersTemplate))
	case types.EmailAddressesTemplate:
		return template.Must(template.New("outer").Funcs(funcMap).Parse(EmailsTemplate))
	case types.CallsTemplate:
		return template.Must(template.New("outer").Funcs(funcMap).Parse(MarkDownTemplateCalls))
	case types.MessagesTemplate:
		return template.Must(template.New("outer").Funcs(funcMap).Parse(MarkDownTemplateMessages))
	case types.NotesSheetTemplate:
		return template.Must(template.New("outer").Funcs(funcMap).Parse(NotesSheetTemplate))
	}

	return nil
}

func (mdData *MarkdownData) GetTemplateData(id types.TemplateIdentifier) []byte {
	switch id {
	case types.ContactSheetTemplate:
		return []byte(ContactSheetTemplate)
	case types.AddressesTemplate:
		return []byte(AddressesTemplate)
	case types.PersonalDataTemplate:
		return []byte(PersonalDataTemplate)
	case types.PhoneNumbersTemplate:
		return []byte(PhoneNumbersTemplate)
	case types.EmailAddressesTemplate:
		return []byte(EmailsTemplate)
	case types.CallsTemplate:
		return []byte(MarkDownTemplateCalls)
	case types.MessagesTemplate:
		return []byte(MarkDownTemplateMessages)
	}

	return nil
}
