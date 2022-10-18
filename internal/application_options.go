package internal

import (
	"errors"
	"os"
)

/*
	memberShipsAsTag  string
	pathToContacts    string
	pathToGroups      string
	templateDirectory string
	pathForFiles      string
*/

func WithSmsBackupFile(smsBackupFile string) ApplicationOption {
	return func(application *Application) error {
		if smsBackupFile == "" {
			return os.ErrNotExist
		}
		fi, err := os.Stat(smsBackupFile)
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return errors.New("specified path is a directory")
		}
		application.smsBackupFile = smsBackupFile
		return nil
	}
}

func WithPathForFiles(pathForFiles string) ApplicationOption {
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

func WithPathToGroups(pathToGroups string) ApplicationOption {
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

func WithPathToContacts(pathToContacts string) ApplicationOption {
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

func WithMembershipsAsTag(memberShipsAsTag string) ApplicationOption {
	return func(application *Application) error {
		application.memberShipsAsTag = memberShipsAsTag
		return nil
	}
}

func WithTemplateDirectory(templateDirectory string) ApplicationOption {
	return func(application *Application) error {
		application.templateDirectory = templateDirectory
		return nil
	}
}
