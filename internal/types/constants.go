package types

const (
	OuterTemplate          TemplateIdentifier = iota
	AddressesTemplate      TemplateIdentifier = iota + 1
	PersonalDataTemplate   TemplateIdentifier = iota + 2
	PhoneNumbersTemplate   TemplateIdentifier = iota + 3
	EmailAddressesTemplate TemplateIdentifier = iota + 4
	CallsTemplate          TemplateIdentifier = iota + 5
	MessagesTemplate       TemplateIdentifier = iota + 6
)
