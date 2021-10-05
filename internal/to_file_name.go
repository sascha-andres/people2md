package internal

import "strings"

func toFileName(in string) string {
	return strings.ReplaceAll(in, ":", " ")
}
