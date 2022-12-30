package markdown

import (
	"text/template"

	"github.com/sascha-andres/people2md/internal/types"
)

type (
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
	switch id {
	case types.OuterTemplate:
		return template.Must(template.New("outer").Parse(MarkDownTemplate))
	case types.AddressesTemplate:
		return template.Must(template.New("outer").Parse(AddressesTemplate))
	case types.PersonalDataTemplate:
		return template.Must(template.New("outer").Parse(PersonalDataTemplate))
	case types.PhoneNumbersTemplate:
		return template.Must(template.New("outer").Parse(PhoneNumbersTemplate))
	case types.EmailAddressesTemplate:
		return template.Must(template.New("outer").Parse(EmailsTemplate))
	case types.CallsTemplate:
		return template.Must(template.New("outer").Parse(MarkDownTemplateCalls))
	case types.MessagesTemplate:
		return template.Must(template.New("outer").Parse(MarkDownTemplateMessages))
	}

	return nil
}

func (mdData *MarkdownData) GetTemplateData(id types.TemplateIdentifier) []byte {
	switch id {
	case types.OuterTemplate:
		return []byte(MarkDownTemplate)
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
