package internal

import (
	"errors"
	"os"
)

func WithDryRun(dryRun bool) Option {
	return func(application *Application) error {
		application.dryRun = dryRun
		return nil
	}
}

func WithFileExtension(fileExtension string) Option {
	return func(application *Application) error {
		application.fileExtension = fileExtension
		return nil
	}
}

func WithCollectionFile(file string) Option {
	return func(application *Application) error {
		application.collectionFile = file
		return nil
	}
}

func WithTagPrefix(prefix string) Option {
	return func(application *Application) error {
		application.tagPrefix = prefix
		return nil
	}
}

func WithVerbose(verbose bool) Option {
	return func(application *Application) error {
		application.verbose = verbose
		return nil
	}
}

func WithPathForFiles(pathForFiles string) Option {
	return func(application *Application) error {
		if pathForFiles == "" {
			return os.ErrNotExist
		}
		fi, err := os.Stat(pathForFiles)
		if err != nil {
			return err
		}
		if !fi.IsDir() {
			return errors.New("specified path is not a directory")
		}
		application.pathForFiles = pathForFiles
		return nil
	}
}

func WithPathToGroups(pathToGroups string) Option {
	return func(application *Application) error {
		if pathToGroups == "" {
			return os.ErrNotExist
		}
		fi, err := os.Stat(pathToGroups)
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return errors.New("specified groups file is a directory")
		}
		application.pathToGroups = pathToGroups
		return nil
	}
}

func WithPathToContacts(pathToContacts string) Option {
	return func(application *Application) error {
		if pathToContacts == "" {
			return os.ErrNotExist
		}
		fi, err := os.Stat(pathToContacts)
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return errors.New("specified contacts file is a directory")
		}
		application.pathToContacts = pathToContacts
		return nil
	}
}

func WithMembershipsAsTag(memberShipsAsTag string) Option {
	return func(application *Application) error {
		application.memberShipsAsTag = memberShipsAsTag
		return nil
	}
}

func WithTemplateDirectory(templateDirectory string) Option {
	return func(application *Application) error {
		application.templateDirectory = templateDirectory
		return nil
	}
}

func WithTemplateGroup(templateGroup string) Option {
	return func(application *Application) error {
		application.templateGroup = templateGroup
		return nil
	}
}
