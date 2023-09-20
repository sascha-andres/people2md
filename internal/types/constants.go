package types

const (
	ContactSheetTemplate   TemplateIdentifier = iota
	AddressesTemplate      TemplateIdentifier = iota * 2
	PersonalDataTemplate   TemplateIdentifier = iota * 2
	PhoneNumbersTemplate   TemplateIdentifier = iota * 2
	EmailAddressesTemplate TemplateIdentifier = iota * 2
	CallsTemplate          TemplateIdentifier = iota * 2
	MessagesTemplate       TemplateIdentifier = iota * 2
	NotesSheetTemplate     TemplateIdentifier = iota * 2
)

var TemplateNames = map[TemplateIdentifier]string{
	ContactSheetTemplate:   "contactsheet",
	NotesSheetTemplate:     "notes",
	AddressesTemplate:      "addresses",
	PersonalDataTemplate:   "personal",
	PhoneNumbersTemplate:   "phone",
	EmailAddressesTemplate: "emails",
	CallsTemplate:          "calls",
	MessagesTemplate:       "messages",
}

var TemplateTypes = map[string]TemplateIdentifier{
	"contactsheet": ContactSheetTemplate,
	"notes":        NotesSheetTemplate,
	"addresses":    AddressesTemplate,
	"personal":     PersonalDataTemplate,
	"phone":        PhoneNumbersTemplate,
	"emails":       EmailAddressesTemplate,
	"calls":        CallsTemplate,
	"messages":     MessagesTemplate,
}

var TemplateFileExtension = ".tmpl"
