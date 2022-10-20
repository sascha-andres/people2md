package internal

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

func (mdData *MarkdownData) WriteMarkdown(pathForFiles string, outer *template.Template, c *Contact) {
	var buff bytes.Buffer
	outer.Execute(&buff, mdData)
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

	err = ioutil.WriteFile(destinationPath, buff.Bytes(), 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not write file: %s", err)
		return
	}
}
