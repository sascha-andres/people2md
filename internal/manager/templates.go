package manager

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/sascha-andres/people2md/internal/generator"
	"github.com/sascha-andres/people2md/internal/types"
)

// ManageTemplates manages the available templates
func ManageTemplates(templateDirectory, templateType, templateGroup string, verbs []string) error {
	if len(verbs) == 0 {
		log.Print("command not provided")
		return nil
	}
	if templateDirectory == "" {
		templateDirectory = "."
	}
	switch verbs[0] {
	case "dump":
		log.Print("dumping templates")
		g, err := generator.GetGenerator()
		if err != nil {
			return err
		}
		t := &types.Templates{Directory: templateDirectory, Group: templateGroup}
		return t.WriteTemplates(g)
	case "list":
		infos, err := os.ReadDir(templateDirectory)
		if err != nil {
			return err
		}
		log.Print("template groups:")
		for _, info := range infos {
			if info.IsDir() {
				if strings.HasPrefix(info.Name(), "template_") {
					groupName, _ := strings.CutPrefix(info.Name(), "template_")
					log.Printf("- %s", groupName)
				}
			}
		}
		return nil
	case "edit":
		if templateType == "" {
			return errors.New("template type not provided")
		}
		if templateGroup == "" {
			return errors.New("you have to provide the template group name")
		}
		if _, ok := types.TemplateTypes[templateType]; !ok {
			return errors.New("unknown template type")
		}
		editor := os.Getenv("EDITOR")
		if editor == "" {
			return errors.New("no editor defined")
		}
		command := exec.Command(editor, path.Join(path.Join(templateDirectory, "template_"+templateGroup), templateType+types.TemplateFileExtension))
		command.Stdout = os.Stdout
		command.Stdin = os.Stdin
		command.Stderr = os.Stderr
		var err error
		if err = command.Start(); err != nil {
			return fmt.Errorf("could not start command: %w", err)
		}
		return command.Wait()
	}
	return errors.New("unknown command")
}
