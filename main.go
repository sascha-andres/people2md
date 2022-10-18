package main

import (
	"log"
	"os"

	"github.com/sascha-andres/flag"
	"livingit.de/code/people2md/internal"
)

var (
	memberShipsAsTag  string
	pathToContacts    string
	pathToGroups      string
	pathForFiles      string
	templateDirectory string
	smsBackupFile     string
)

func init() {
	log.SetPrefix("P2MD")
	log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)

	flag.SetEnvPrefix("P2MD")
	flag.StringVar(&memberShipsAsTag, "tags", "", "list of labels to convert to tags")
	flag.StringVar(&pathToContacts, "contacts", "contacts.json", "output of goobook dump_contacts")
	flag.StringVar(&pathToGroups, "groups", "groups.json", "output of goobook dump_groups")
	flag.StringVar(&pathForFiles, "output", ".", "store output in this directory")
	flag.StringVar(&templateDirectory, "template-directory", "", "load templates from directcory")
	flag.StringVar(&smsBackupFile, "sms", "", "path to sms backup file")
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

	app, err := internal.NewApplication(
		internal.WithPathForFiles(pathForFiles),
		internal.WithMembershipsAsTag(memberShipsAsTag),
		internal.WithPathToContacts(pathToContacts),
		internal.WithSmsBackupFile(smsBackupFile),
		internal.WithPathToGroups(pathToGroups),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
