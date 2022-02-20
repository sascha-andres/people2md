package internal

import "os"

// WriteTemplates writes default templates to current directory
func (t *Templates) WriteTemplates() error {
	if err := os.WriteFile("addresses.tmpl", []byte(AddressesTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile("emails.tmpl", []byte(EmailsTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile("markdown.tmpl", []byte(MarkDownTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile("personal.tmpl", []byte(PersonalDataTemplate), 0600); err != nil {
		return err
	}
	return os.WriteFile("phone.tmpl", []byte(PhoneNumbersTemplate), 0600)
}
