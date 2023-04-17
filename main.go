package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sascha-andres/people2md/internal/generator"

	"github.com/sascha-andres/people2md/internal"
	"github.com/sascha-andres/reuse/flag"
)

var (
	collectionFile    string
	memberShipsAsTag  string
	pathToContacts    string
	pathToGroups      string
	pathForFiles      string
	templateDirectory string
	smsBackupFile     string
	callBackupFile    string
	tagPrefix         string
	fileExtension     string
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
	flag.StringVar(&collectionFile, "collection-file", "", "get calls and messages from collection")
	flag.StringVar(&smsBackupFile, "sms", "", "path to sms backup file")
	flag.StringVar(&callBackupFile, "calls", "", "path to call backup file")
	flag.BoolVar(&verbose, "verbose", false, "print some output while operating")
	flag.StringVar(&tagPrefix, "tag-prefix", "", "prefix tag with string")
	flag.StringVar(&fileExtension, "extension", "md", "file extension for output files")
}

func run() error {
	arguments := os.Args

	verbs := flag.GetVerbs()
	fmt.Printf("%#v\n", verbs)

	if len(arguments) >= 2 {
		if arguments[1] == "help" {
			flag.PrintDefaults()
			return nil
		}
		if arguments[1] == "dump-templates" {
			g, err := generator.GetGenerator()
			if err != nil {
				log.Fatal(err)
			}
			if err = internal.WriteTemplates(g); err != nil {
				log.Fatal(err)
			}
			return nil
		}
	}

	return nil

	app, err := internal.NewApplication(
		internal.WithPathForFiles(pathForFiles),
		internal.WithMembershipsAsTag(memberShipsAsTag),
		internal.WithPathToContacts(pathToContacts),
		internal.WithSmsBackupFile(smsBackupFile),
		internal.WithPathToGroups(pathToGroups),
		internal.WithCallBackupFile(callBackupFile),
		internal.WithVerbose(verbose),
		internal.WithTagPrefix(tagPrefix),
		internal.WithCollectionFile(collectionFile),
		internal.WithFileExtension(fileExtension),
	)

	if err != nil {
		return err
	}

	return app.Run()
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
