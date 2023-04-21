package main

import (
	"log"

	"github.com/sascha-andres/people2md/internal"
	"github.com/sascha-andres/people2md/internal/manager"
	"github.com/sascha-andres/reuse/flag"
)

var (
	collectionFile, callBackupFile, smsBackupFile, pathToGroups, pathToContacts string
	memberShipsAsTag, tagPrefix                                                 string
	verbose                                                                     bool
	pathForFiles, fileExtension                                                 string
	templateGroup, templateType, templateDirectory                              string
)

func init() {
	log.SetPrefix("[P2MD] ")
	log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)

	flag.SetEnvPrefix("P2MD")
	// template specific flags
	flag.StringVar(&templateDirectory, "template-directory", "", "load templates from directory")
	flag.StringVar(&templateType, "template-type", "", "used to edit templates (eg the phone number sub template)")
	flag.StringVar(&templateGroup, "template-group", "md", "used to edit templates (eg the phone number sub template)")
	flag.StringVar(&memberShipsAsTag, "tags", "", "list of labels to convert to tags")
	flag.StringVar(&tagPrefix, "tag-prefix", "", "prefix tag with string")
	// data handling flags
	flag.StringVar(&pathToContacts, "contacts", "contacts.json", "output of goobook dump_contacts")
	flag.StringVar(&pathToGroups, "groups", "groups.json", "output of goobook dump_groups")
	flag.StringVar(&collectionFile, "collection-file", "", "get calls and messages from collection")
	flag.StringVar(&smsBackupFile, "sms", "", "path to sms backup file")
	flag.StringVar(&callBackupFile, "calls", "", "path to call backup file")
	// output flags
	flag.StringVar(&pathForFiles, "output", ".", "store output in this directory")
	flag.BoolVar(&verbose, "verbose", false, "print some output while operating")
	flag.StringVar(&fileExtension, "extension", "md", "file extension for output files")
}

func run() error {
	verbs := flag.GetVerbs()

	if len(verbs) > 0 {
		if verbs[0] == "help" {
			flag.PrintDefaults()
			return nil
		}
		if verbs[0] == "templates" {
			return manager.ManageTemplates(templateDirectory, templateType, templateGroup, verbs[1:])
		}
	}

	app, err := internal.NewApplication(
		// template specific arguments
		internal.WithTemplateDirectory(templateDirectory),
		internal.WithTemplateGroup(templateGroup),
		internal.WithMembershipsAsTag(memberShipsAsTag),
		internal.WithTagPrefix(tagPrefix),
		// data handling arguments
		internal.WithPathToContacts(pathToContacts),
		internal.WithSmsBackupFile(smsBackupFile),
		internal.WithPathToGroups(pathToGroups),
		internal.WithCallBackupFile(callBackupFile),
		internal.WithCollectionFile(collectionFile),
		// output arguments
		internal.WithPathForFiles(pathForFiles),
		internal.WithVerbose(verbose),
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
