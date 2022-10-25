package internal

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/sascha-andres/people2md/internal/types"
	"github.com/sascha-andres/sbrdata"
	"log"
	"os"
	"path"
	"text/template"
)

func Handle(c *types.Contact, generator types.DataBuilder, pathForFiles, tags string, personalData *template.Template, groups []types.ContactGroup, addresses, phoneNumbers, emailAddresses, outer *template.Template, sms sbrdata.Messages, calls sbrdata.Calls, verbose bool) {
	e := &types.Elements{}

	e.ETag = c.Etag
	e.ResourceName = c.ResourceName

	generator.SetETag(c.Etag)
	generator.SetResourceName(c.ResourceName)

	e.Calls = generator.BuildCalls(calls, c)
	e.Sms = generator.BuildSms(sms, c)
	e.PersonalData = generator.BuildPersonalData(personalData, c)
	e.Tags = generator.BuildTags(tags, c, groups)
	e.Addresses = generator.BuildAddresses(c, addresses)
	e.PhoneNumbers = generator.BuildPhoneNumbers(c, phoneNumbers)
	e.Email = generator.BuildEmailAddresses(c, emailAddresses)

	var buff bytes.Buffer
	outer.Execute(&buff, e)
	var fileName = ""
	if len(c.Names) > 0 {
		fileName = toFileName(c.Names[0].DisplayName) + ".md"
	} else {
		fileName = toFileName(c.Organizations[0].Name) + ".md"
	}

	destinationPath := path.Join(pathForFiles, fileName)

	if _, err := os.Stat(destinationPath); errors.Is(err, os.ErrNotExist) {
		os.WriteFile(destinationPath, buff.Bytes(), 0600)
		return
	}

	hasher := sha256.New()
	hasher.Write(buff.Bytes())
	hashNew := hasher.Sum(nil)

	fileData, err := os.ReadFile(destinationPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read existing file: %s", err)
		return
	}

	hasher = sha256.New()
	hasher.Write(fileData)
	hashOld := hasher.Sum(nil)

	res := bytes.Compare(hashOld, hashNew)

	if res == 0 {
		return
	}

	if verbose {
		log.Printf("replacing %s", destinationPath)
	}
	err = os.WriteFile(destinationPath, buff.Bytes(), 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not write file: %s", err)
		return
	}
}
