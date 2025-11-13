package markdown

import (
	"fmt"
	"strings"

	"github.com/sascha-andres/people2md/internal/types"
)

// BuildAliases builds the aliases for a contact
func (mdData *MarkdownData) BuildAliases(c *types.Contact) []string {
	var aliases []string
	if len(c.Names) > 0 {
		for _, name := range c.Names {
			a := buildAlias(name.DisplayName)
			if a == "" {
				continue
			}
			aliases = append(aliases, a)
		}
	}
	return aliases
}

var replacementRunes = map[rune]string{
	'è': "e",
	'é': "e",
	'ê': "e",
	'ë': "e",
	'à': "a",
	'á': "a",
	'â': "a",
	'ä': "a",
	'ò': "o",
	'ó': "o",
	'ô': "o",
	'ö': "o",
	'ì': "i",
	'í': "i",
	'î': "i",
	'ï': "i",
	'ù': "u",
	'ú': "u",
	'û': "u",
	'ü': "u",
	'ñ': "n",
	'ß': "ss",
}

// buildAlias builds an alias from a name
func buildAlias(name string) string {
	result := ""
	for from, to := range replacementRunes {
		if result == "" {
			if strings.Contains(name, string(from)) {
				result = strings.Replace(name, string(from), to, -1)
			}
		} else {
			result = strings.Replace(result, string(from), to, -1)
		}
	}
	if strings.Contains(result, "Carmen") {
		fmt.Printf("Carmen: %s\n", result)
	}
	return result
}
