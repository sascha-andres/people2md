package internal

import "strings"

var replacements = map[string]string{
	" ": " ",
	//"ä": "ae",
	//"ö": "oe",
	//"ü": "ue",
	//"Ä": "Ae",
	//"Ö": "Oe",
	//"Ü": "Ue",
	"ß": "ss",
	":": "",
}

// toFileName converts a string to a valid file name.
func toFileName(in string) string {
	for k, v := range replacements {
		in = strings.ReplaceAll(in, k, v)
	}
	return strings.TrimSpace(in)
}
