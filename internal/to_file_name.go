package internal

import "strings"

func toFileName(in string) string {
	return strings.TrimSpace(strings.ReplaceAll(in, ":", " "))
}
