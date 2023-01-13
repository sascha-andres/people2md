package main

import (
	"log"
	"os"

	"github.com/sascha-andres/people2md/internal/generator"

	"github.com/sascha-andres/flag"
	"github.com/sascha-andres/people2md/internal"
)

var (
	memberShipsAsTag  string
	pathToContacts    string
	pathToGroups      string
	pathForFiles      string
	templateDirectory string
	smsBackupFile     string
	callBackupFile    string
	tagPrefix         string
	verbose           bool
)

func init() {
	log.SetPrefix("[P2MD] ")
	log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)

	flag.SetEnvPrefix("P2MD")
	flag.StringVar(&memberShipsAsTag, "tags", "", "list of labels to convert to tags")
	flag.StringVar(&pathToContacts, "contacts", "contacts.json", "output of goobook dump_contacts")
	flag.StringVar(&pathToGroups, "groups", "groups.json", "output of goobook dump_groups")
	flag.StringVar(&pathForFiles, "output", ".", "store output in this directory")
	flag.StringVar(&templateDirectory, "template-directory", "", "load templates from directcory")
	flag.StringVar(&smsBackupFile, "sms", "", "path to sms backup file")
	flag.StringVar(&callBackupFile, "calls", "", "path to call backup file")
	flag.BoolVar(&verbose, "verbose", false, "print some output while operating")
	flag.StringVar(&tagPrefix, "tag-prefix", "", "prefix tag with string")
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
			g, err := generator.GetGenerator()
			if err != nil {
				log.Fatal(err)
			}
			if err := internal.WriteTemplates(g); err != nil {
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
		internal.WithCallBackupFile(callBackupFile),
		internal.WithVerbose(verbose),
		internal.WithTagPrefix(tagPrefix),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
