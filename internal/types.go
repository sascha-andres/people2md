package internal

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"strings"

	"github.com/sascha-andres/people2md/internal/generator"
	"github.com/sascha-andres/people2md/internal/types"
	"github.com/sascha-andres/sbrdata"
)

type (

	// Application is the root of the functionality except some infrastructure stuff
	Application struct {
		memberShipsAsTag  string
		pathToContacts    string
		pathToGroups      string
		templateDirectory string
		pathForFiles      string
		smsBackupFile     string
		callBackupFile    string
		verbose           bool
		tagPrefix         string
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
	var contacts []types.Contact
	log.Print("reading contacts")
	if err := json.Unmarshal(data, &contacts); err != nil {
		return err
	}

	log.Print("reading groups")
	data, err = os.ReadFile(app.pathToGroups)
	if err != nil {
		return err
	}
	var groups []types.ContactGroup
	if err := json.Unmarshal(data, &groups); err != nil {
		return err
	}

	var sms sbrdata.Messages
	if app.smsBackupFile != "" {
		log.Print("reading SMS backup file")
		data, err := os.ReadFile(app.smsBackupFile)
		if err != nil {
			return err
		}
		err = xml.Unmarshal(data, &sms)
		if err != nil {
			return err
		}
	}

	var calls sbrdata.Calls
	if app.callBackupFile != "" {
		log.Print("reading call backup file")
		data, err := os.ReadFile(app.callBackupFile)
		if err != nil {
			return err
		}
		err = xml.Unmarshal(data, &calls)
		if err != nil {
			return err
		}
	}

	g, err := generator.GetGenerator()
	if err != nil {
		return err
	}

	templates, err := NewTemplates(g, app.templateDirectory)
	if err != nil {
		return err
	}

	for _, c := range contacts {
		if 0 == len(c.Names) && 0 == len(c.Organizations) {
			continue
		}
		if len(c.Names) > 0 {
			c.Names[0].DisplayName = strings.TrimSpace(c.Names[0].DisplayName)
		}
		if len(c.Organizations) > 0 {
			c.Organizations[0].Name = strings.TrimSpace(c.Organizations[0].Name)
		}
		Handle(&c, g, app.pathForFiles,
			app.memberShipsAsTag,
			app.tagPrefix,
			templates.PersonalData,
			groups,
			templates.Addresses,
			templates.PhoneNumbers,
			templates.EmailAddresses,
			templates.Outer,
			sms,
			calls,
			app.verbose,
			templates.Calls,
			templates.Messages,
		)
	}
	return nil
}
