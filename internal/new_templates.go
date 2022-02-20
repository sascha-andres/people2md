package internal

import "text/template"

// NewTemplates returns a new instance of the templates struct containing
// the templates used to render the contacts
func NewTemplates(directory string) *Templates {
	return &Templates{
		Outer:          template.Must(template.New("outer").Parse(MarkDownTemplate)),
		Addresses:      template.Must(template.New("addresses").Parse(AddressesTemplate)),
		PersonalData:   template.Must(template.New("personalData").Parse(PersonalDataTemplate)),
		PhoneNumbers:   template.Must(template.New("phoneNumbers").Parse(PhoneNumbersTemplate)),
		EmailAddresses: template.Must(template.New("emailAddresses").Parse(EmailsTemplate)),
	}
}
