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
	memberShipsAsTag string
	pathToContacts   string
	pathToGroups     string
	pathForFiles     string
)

func init() {
	flag.StringVar(&memberShipsAsTag, "tags", "", "list of labels to convert to tags")
	flag.StringVar(&pathToContacts, "contacts", "contacts.json", "output of goobook dump_contacts")
	flag.StringVar(&pathToGroups, "groups", "groups.json", "output of goobook dump_groups")
	flag.StringVar(&pathForFiles, "output", ".", "store output in this directory")
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
			return
		}
	}

	data, err := ioutil.ReadFile(pathToContacts)
	if err != nil {
		log.Fatal(err)
	}
	var contacts []internal.Contact
	if err := json.Unmarshal([]byte(data), &contacts); err != nil {
		log.Fatal(err)
	}

	data, err = ioutil.ReadFile(pathToGroups)
	if err != nil {
		log.Fatal(err)
	}
	var groups []internal.ContactGroup
	if err := json.Unmarshal([]byte(data), &groups); err != nil {
		log.Fatal(err)
	}

	templates := internal.NewTemplates("")

	for _, c := range contacts {
		if 0 == len(c.Names) && 0 == len(c.Organizations) {
			continue
		}
		c.Handle(pathForFiles, memberShipsAsTag, templates.PersonalData, groups, templates.Addresses, templates.PhoneNumbers, templates.EmailAddresses, templates.Outer)
	}
}
