package types

import "strings"

// TemplateReplace is used to replace a string in a template
func TemplateReplace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}
