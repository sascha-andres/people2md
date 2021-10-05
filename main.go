package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"livingit.de/code/people2md/internal"
	"log"
	"text/template"
)

var (
	memberShipsAsTag string
	pathToContacts   string
	pathToGroups   string
)

func init() {
	flag.StringVar(&memberShipsAsTag, "tags", "", "list of labels to convert to tags")
	flag.StringVar(&pathToContacts, "contacts", "contacts.json", "output of goobook dump_contacts")
	flag.StringVar(&pathToGroups, "groups", "groups.json", "output of goobook dump_groups")
}

func main() {
	flag.Parse()

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

	outer := template.Must(template.New("outer").Parse(internal.MarkDownTemplate))
	addresses := template.Must(template.New("addresses").Parse(internal.AddressesTemplate))
	personalData := template.Must(template.New("personalData").Parse(internal.PersonalDataTemplate))
	phoneNumbers := template.Must(template.New("phoneNumbers").Parse(internal.PhoneNumbersTemplate))
	emailAddresses := template.Must(template.New("emailAddresses").Parse(internal.EmailsTemplate))

	for _, c := range contacts {
		if 0 == len(c.Names) && 0 == len(c.Organizations) {
			continue
		}
		c.Handle(memberShipsAsTag, personalData, groups, addresses, phoneNumbers, emailAddresses, outer)
	}
}
