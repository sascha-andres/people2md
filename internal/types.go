package internal

import (
	"encoding/json"
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
		pathForFiles      string
		verbose           bool
		tagPrefix         string
		collectionFile    string
		fileExtension     string
		templateDirectory string
		templateGroup     string
		dryRun            bool
	}

	Option func(application *Application) error
)

// NewApplication returns the app root
func NewApplication(opts ...Option) (*Application, error) {
	a := &Application{}
	for i := range opts {
		err := opts[i](a)
		if err != nil {
			return nil, err
		}
	}
	if a.collectionFile == "" {
		return nil, errors.New("no comm data provided")
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

	//var sms *sbrdata.Messages
	//var calls *sbrdata.Calls

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

	templates := &types.Templates{Directory: app.templateDirectory, Group: app.templateGroup}
	err = templates.NewTemplates(g)
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
			Collection:   collection,
		}, g, templates)
	}
	return nil
}
