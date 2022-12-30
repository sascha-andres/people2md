package internal

import (
	"os"

	"github.com/sascha-andres/people2md/internal/types"
)

// WriteTemplates writes default templates to current directory
func WriteTemplates(db types.DataBuilder) error {
	if err := os.WriteFile("addresses.tmpl", db.GetTemplateData(types.AddressesTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile("emails.tmpl", db.GetTemplateData(types.EmailAddressesTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile("markdown.tmpl", db.GetTemplateData(types.OuterTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile("personal.tmpl", db.GetTemplateData(types.PersonalDataTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile("calls.tmpl", db.GetTemplateData(types.CallsTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile("messages.tmpl", db.GetTemplateData(types.MessagesTemplate), 0600); err != nil {
		return err
	}
	return os.WriteFile("phone.tmpl", db.GetTemplateData(types.PhoneNumbersTemplate), 0600)
}
