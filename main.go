package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"livingit.de/code/people2md/internal"
	"log"
	"os"
)

var (
	memberShipsAsTag  string
	pathToContacts    string
	pathToGroups      string
	pathForFiles      string
	templateDirectory string
)

func init() {
	flag.StringVar(&memberShipsAsTag, "tags", "", "list of labels to convert to tags")
	flag.StringVar(&pathToContacts, "contacts", "contacts.json", "output of goobook dump_contacts")
	flag.StringVar(&pathToGroups, "groups", "groups.json", "output of goobook dump_groups")
	flag.StringVar(&pathForFiles, "output", ".", "store output in this directory")
	flag.StringVar(&templateDirectory, "template-directory", "", "load templates from directcory")
}

func main() {
	flag.Parse()

	arguments := os.Args
	if len(arguments) >= 2 {
		if arguments[1] == "help" {
			flag.PrintDefaults()
			return
		}
		if arguments[1] == "dump-templates" {
			t, err := internal.NewTemplates("")
			if err != nil {
				log.Fatal(err)
			}
			if err := t.WriteTemplates(); err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	data, err := ioutil.ReadFile(pathToContacts)
	if err != nil {
		log.Fatal(err)
	}
	var contacts []internal.Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		log.Fatal(err)
	}

	data, err = ioutil.ReadFile(pathToGroups)
	if err != nil {
		log.Fatal(err)
	}
	var groups []internal.ContactGroup
	if err := json.Unmarshal(data, &groups); err != nil {
		log.Fatal(err)
	}

	templates, err := internal.NewTemplates(templateDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range contacts {
		if 0 == len(c.Names) && 0 == len(c.Organizations) {
			continue
		}
		c.Handle(pathForFiles,
			memberShipsAsTag,
			templates.PersonalData,
			groups,
			templates.Addresses,
			templates.PhoneNumbers,
			templates.EmailAddresses,
			templates.Outer,
		)
	}
}
