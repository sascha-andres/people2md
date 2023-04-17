package internal

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"os"

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
		collectionFile    string
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
	if a.collectionFile != "" && (a.smsBackupFile != "" || a.callBackupFile != "") {
		return nil, errors.New("either use collection of single files")
	}
	return a, nil
}

type Contacts []types.Contact
type Groups []types.ContactGroup

func (app *Application) getContacts() (Contacts, error) {
	data, err := os.ReadFile(app.pathToContacts)
	if err != nil {
		return nil, err
	}
	var contacts []types.Contact
	log.Print("reading contacts")
	if err := json.Unmarshal(data, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}

func (app *Application) getGroups() (Groups, error) {
	log.Print("reading groups")
	data, err := os.ReadFile(app.pathToGroups)
	if err != nil {
		return nil, err
	}
	var groups []types.ContactGroup
	if err := json.Unmarshal(data, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

// Run executes the application
func (app *Application) Run() error {
	contacts, err := app.getContacts()
	if err != nil {
		return err
	}
	groups, err := app.getGroups()
	if err != nil {
		return err
	}

	var sms *sbrdata.Messages
	sms, err = app.getMessages()
	if err != nil {
		return err
	}

	var calls *sbrdata.Calls
	calls, err = app.getCalls()
	if err != nil {
		return err
	}

	var collection *sbrdata.Collection
	if app.collectionFile != "" {
		collection, err = sbrdata.LoadCollection(app.collectionFile) // TODO this should fail if it does not exist
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
		app.handle(types.DataReferences{
			Contact:      &c,
			PathForFiles: app.pathForFiles,
			Tags:         app.memberShipsAsTag,
			TagPrefix:    app.tagPrefix,
			Groups:       groups,
			Sms:          sms,
			CallData:     calls,
			Collection:   collection,
		}, g, templates)
	}
	return nil
}

func (app *Application) getCalls() (*sbrdata.Calls, error) {
	var calls sbrdata.Calls
	if app.callBackupFile != "" {
		log.Print("reading call backup file")
		data, err := os.ReadFile(app.callBackupFile)
		if err != nil {
			return nil, err
		}
		err = xml.Unmarshal(data, &calls)
		if err != nil {
			return nil, err
		}
	}
	return &calls, nil
}

func (app *Application) getMessages() (*sbrdata.Messages, error) {
	var sms sbrdata.Messages
	if app.smsBackupFile != "" {
		log.Print("reading SMS backup file")
		data, err := os.ReadFile(app.smsBackupFile)
		if err != nil {
			return nil, err
		}
		err = xml.Unmarshal(data, &sms)
		if err != nil {
			return nil, err
		}
	}
	return &sms, nil
}
