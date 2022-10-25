package internal

import (
	"github.com/sascha-andres/people2md/internal/types"
	"os"
	"path"
	"text/template"
)

// NewTemplates returns a new instance of the templates struct containing
// the templates used to render the contacts
func NewTemplates(generator types.DataBuilder, directory string) (*types.Templates, error) {
	t := &types.Templates{
		Outer:          generator.GetTemplate(types.OuterTemplate),
		Addresses:      generator.GetTemplate(types.AddressesTemplate),
		PersonalData:   generator.GetTemplate(types.PersonalDataTemplate),
		PhoneNumbers:   generator.GetTemplate(types.PhoneNumbersTemplate),
		EmailAddresses: generator.GetTemplate(types.EmailAddressesTemplate),
	}
	if "" != directory {
		if fileExists(path.Join(directory, "addresses.tmpl")) {
			data, err := os.ReadFile(path.Join(directory, "addresses.tmpl"))
			if err != nil {
				return nil, err
			}
			t.Addresses = template.Must(template.New("addresses").Parse(string(data)))
		}
		if fileExists(path.Join(directory, "emails.tmpl")) {
			data, err := os.ReadFile(path.Join(directory, "emails.tmpl"))
			if err != nil {
				return nil, err
			}
			t.EmailAddresses = template.Must(template.New("emailAddresses").Parse(string(data)))
		}
		if fileExists(path.Join(directory, "markdown.tmpl")) {
			data, err := os.ReadFile(path.Join(directory, "markdown.tmpl"))
			if err != nil {
				return nil, err
			}
			t.Outer = template.Must(template.New("outer").Parse(string(data)))
		}
		if fileExists(path.Join(directory, "personal.tmpl")) {
			data, err := os.ReadFile(path.Join(directory, "personal.tmpl"))
			if err != nil {
				return nil, err
			}
			t.PersonalData = template.Must(template.New("personalData").Parse(string(data)))
		}
		if fileExists(path.Join(directory, "phone.tmpl")) {
			data, err := os.ReadFile(path.Join(directory, "phone.tmpl"))
			if err != nil {
				return nil, err
			}
			t.PhoneNumbers = template.Must(template.New("phoneNumbers").Parse(string(data)))
		}
	}
	return t, nil
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}
