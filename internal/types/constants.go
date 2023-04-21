package types

const (
	ContactSheetTemplate   TemplateIdentifier = iota
	AddressesTemplate      TemplateIdentifier = iota + 1
	PersonalDataTemplate   TemplateIdentifier = iota + 2
	PhoneNumbersTemplate   TemplateIdentifier = iota + 3
	EmailAddressesTemplate TemplateIdentifier = iota + 4
	CallsTemplate          TemplateIdentifier = iota + 5
	MessagesTemplate       TemplateIdentifier = iota + 6
)

var TemplateNames = map[TemplateIdentifier]string{
	ContactSheetTemplate:   "contactsheet",
	AddressesTemplate:      "addresses",
	PersonalDataTemplate:   "personal",
	PhoneNumbersTemplate:   "phone",
	EmailAddressesTemplate: "emails",
	CallsTemplate:          "calls",
	MessagesTemplate:       "messages",
}

var TemplateTypes = map[string]TemplateIdentifier{
	"contactsheet": ContactSheetTemplate,
	"addresses":    AddressesTemplate,
	"personal":     PersonalDataTemplate,
	"phone":        PhoneNumbersTemplate,
	"emails":       EmailAddressesTemplate,
	"calls":        CallsTemplate,
	"messages":     MessagesTemplate,
}

var TemplateFileExtension = ".tmpl"
