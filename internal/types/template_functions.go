package types

import (
	"strings"
	"time"
)

// TemplateReplace is used to replace a string in a template
func TemplateReplace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

// TemplateDate is used to format a date in a template
func TemplateDate(year, month, day uint) string {
	d := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
	return d.Format("2006-01-02")
}
