package internal

import (
	"encoding/json"
	"os"
	"text/template"
)

type (
	Templates struct {
		Outer          *template.Template
		Addresses      *template.Template
		PersonalData   *template.Template
		PhoneNumbers   *template.Template
		EmailAddresses *template.Template
	}

	MarkdownData struct {
		ETag         string
		ResourceName string
		PersonalData string
		Addresses    string
		Im           string
		PhoneNumbers string
		Email        string
		Tags         string
	}

	ContactGroup struct {
		Etag          string
		FormattedName string
		GroupType     string
		MetaData      ContactGroupMetaData
		Name          string
		ResourceName  string
	}

	ContactGroupMetaData struct {
		UpdateTime string
	}

	Contact struct {
		Etag           string
		Memberships    []Membership
		Names          []Name
		PhoneNumbers   []PhoneNumber
		ResourceName   string
		EmailAddresses []EmailAddress
		Organizations  []Organization
		ImClients      []ImClient
		Birthdays      []Birthday
		Addresses      []Address
	}

	Address struct {
		MetaData        *MetaData
		FormattedType   string
		FormattedValue  string
		Type            string
		City            string
		Country         string
		ExtendedAddress string
		PostalCode      string
		StreetAddress   string
	}

	Birthday struct {
		MetaData *MetaData
		Date     *Date
		Text     *string
	}

	Date struct {
		Day   uint
		Month uint
		Year  uint
	}

	ImClient struct {
		FormattedProtocol string
		MetaData          *MetaData
		Protocol          string
		Username          string
	}

	Organization struct {
		FormattedType string
		MetaData      *MetaData
		Name          string
		Type          string
		Title         string
	}

	EmailAddress struct {
		FormattedType string
		MetaData      *MetaData
		Type          string
		Value         string
	}

	PhoneNumber struct {
		CanonicalForm string
		FormattedType string
		MetaData      *MetaData
		Value         string
		Type          string
	}

	Name struct {
		DisplayName          string
		DisplayNameLastFirst string
		FamilyName           string
		GivenName            string
		MetaData             *MetaData
		MiddleName           string
		UnstructuredName     string
	}

	Membership struct {
		ContactGroupMembership *ContactGroupMembership
		MetaData               *MetaData
	}

	ContactGroupMembership struct {
		ContactGroupId           string
		ContactGroupResourceName string
	}

	MetaData struct {
		Primary bool
		Source  *Source
	}

	Source struct {
		Id   string
		Type string
	}

	// Application is the root of the functionality except some infrastructure stuff
	Application struct {
		memberShipsAsTag  string
		pathToContacts    string
		pathToGroups      string
		templateDirectory string
		pathForFiles      string
		smsBackupFile     string
	}

	ApplicationOption func(application *Application) error
)

// NewApplication returns the app root
func NewApplication(opts ...ApplicationOption) (*Application, error) {
	a := &Application{}
	for i := range opts {
		err := opts[i](a)
		if err != nil {
			return nil, err
		}
	}
	return a, nil
}

// Run executes the application
func (app *Application) Run() error {
	data, err := os.ReadFile(app.pathToContacts)
	if err != nil {
		return err
	}
	var contacts []Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		return err
	}

	data, err = os.ReadFile(app.pathToGroups)
	if err != nil {
		return err
	}
	var groups []ContactGroup
	if err := json.Unmarshal(data, &groups); err != nil {
		return err
	}

	templates, err := NewTemplates(app.templateDirectory)
	if err != nil {
		return err
	}

	for _, c := range contacts {
		if 0 == len(c.Names) && 0 == len(c.Organizations) {
			continue
		}
		c.Handle(app.pathForFiles,
			app.memberShipsAsTag,
			templates.PersonalData,
			groups,
			templates.Addresses,
			templates.PhoneNumbers,
			templates.EmailAddresses,
			templates.Outer,
		)
	}
	return nil
}
