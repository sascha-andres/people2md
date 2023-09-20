package types

import (
	"errors"
	"fmt"
	"os"
	"path"
	"text/template"
)

var funcMap = template.FuncMap{
	"replace": TemplateReplace,
}

// loadTemplates loads the templates from the given directory
func loadTemplates(directory string, id TemplateIdentifier) (*template.Template, error) {
	if fileExists(path.Join(directory, TemplateNames[id]+TemplateFileExtension)) {
		data, err := os.ReadFile(path.Join(directory, TemplateNames[id]+TemplateFileExtension))
		if err != nil {
			return nil, err
		}
		return template.Must(template.New(fmt.Sprintf("%d", id)).Funcs(funcMap).Parse(string(data))), nil
	}
	return nil, nil
}

// NewTemplates returns a new instance of the templates struct containing
// the templates used to render the contacts
func (t *Templates) NewTemplates(generator DataBuilder) error {
	t.ContactSheet = generator.GetTemplate(ContactSheetTemplate)
	t.NotesSheet = generator.GetTemplate(NotesSheetTemplate)
	t.Addresses = generator.GetTemplate(AddressesTemplate)
	t.PersonalData = generator.GetTemplate(PersonalDataTemplate)
	t.PhoneNumbers = generator.GetTemplate(PhoneNumbersTemplate)
	t.EmailAddresses = generator.GetTemplate(EmailAddressesTemplate)
	t.Calls = generator.GetTemplate(CallsTemplate)
	t.Messages = generator.GetTemplate(MessagesTemplate)
	if t.Group == "" {
		return nil
	}
	if t.Directory == "" {
		t.Directory = "."
	}
	directory := path.Join(t.Directory, "template_"+t.Group)
	if !directoryExists(directory) {
		return nil
	}
	var err error
	t.Addresses, err = loadTemplates(directory, AddressesTemplate)
	if err != nil {
		return err
	}
	t.EmailAddresses, err = loadTemplates(directory, EmailAddressesTemplate)
	if err != nil {
		return err
	}
	t.ContactSheet, err = loadTemplates(directory, ContactSheetTemplate)
	if err != nil {
		return err
	}
	t.NotesSheet, err = loadTemplates(directory, NotesSheetTemplate)
	if err != nil {
		return err
	}
	t.PersonalData, err = loadTemplates(directory, PersonalDataTemplate)
	if err != nil {
		return err
	}
	t.PhoneNumbers, err = loadTemplates(directory, PhoneNumbersTemplate)
	if err != nil {
		return err
	}
	t.Calls, err = loadTemplates(directory, CallsTemplate)
	if err != nil {
		return err
	}
	t.Messages, err = loadTemplates(directory, MessagesTemplate)
	return err
}

// WriteTemplates writes default templates to current directory
func (t *Templates) WriteTemplates(generator DataBuilder) error {
	if t.Group == "" {
		return errors.New("group is empty")
	}
	if t.Directory == "" {
		t.Directory = "."
	}
	templateDirectory := path.Join(t.Directory, "template_"+t.Group)
	if !fileExists(templateDirectory) {
		if err := os.MkdirAll(templateDirectory, 0700); err != nil {
			return err
		}
		info, err := os.Stat(templateDirectory)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path is not a directory")
		}
	}
	if err := os.WriteFile(path.Join(templateDirectory, TemplateNames[AddressesTemplate]+TemplateFileExtension), generator.GetTemplateData(AddressesTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile(path.Join(templateDirectory, TemplateNames[EmailAddressesTemplate]+TemplateFileExtension), generator.GetTemplateData(EmailAddressesTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile(path.Join(templateDirectory, TemplateNames[ContactSheetTemplate]+TemplateFileExtension), generator.GetTemplateData(ContactSheetTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile(path.Join(templateDirectory, TemplateNames[PersonalDataTemplate]+TemplateFileExtension), generator.GetTemplateData(PersonalDataTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile(path.Join(templateDirectory, TemplateNames[CallsTemplate]+TemplateFileExtension), generator.GetTemplateData(CallsTemplate), 0600); err != nil {
		return err
	}
	if err := os.WriteFile(path.Join(templateDirectory, TemplateNames[MessagesTemplate]+TemplateFileExtension), generator.GetTemplateData(MessagesTemplate), 0600); err != nil {
		return err
	}
	return os.WriteFile(path.Join(templateDirectory, TemplateNames[PhoneNumbersTemplate]+TemplateFileExtension), generator.GetTemplateData(PhoneNumbersTemplate), 0600)
}

func directoryExists(filename string) bool {
	if i, err := os.Stat(filename); err == nil && i.IsDir() {
		return true
	}
	return false
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}
