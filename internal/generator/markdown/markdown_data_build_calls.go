package markdown

import (
	"fmt"
	"strings"

	"github.com/sascha-andres/people2md/internal/types"
	"github.com/sascha-andres/sbrdata"
)

func (mdData *MarkdownData) BuildCalls(calls *sbrdata.Calls, c *types.Contact) string {
	result := ""

	for _, call := range calls.Call {
		include := false
		for _, name := range c.Names {
			include = name.DisplayName == call.ContactName
			if include {
				break
			}
		}
		if !include {
			for _, org := range c.Organizations {
				include = org.Name == call.ContactName
				if include {
					break
				}
			}
		}
		// TODO: check whether it is feasible to
		//  include sth like contains with a cut off of
		//  the last number (having 1234560 as the central
		//  number but identify 12345678 also for this
		//  contact
		if !include && call.Number != "" {
			num := call.Number
			if strings.HasPrefix(call.Number, "0") {
				num = num[1:]
			}
			for _, p := range c.PhoneNumbers {
				include = strings.HasSuffix(p.Value, num)
				if include {
					break
				}
			}
		}
		if include {
			if result == "" {
				result = `|Date|Direction|Number|Duration (s)|
|---|---|---|---|`
			}
			direction := "incoming"
			if call.Type == "2" {
				direction = "outgoing"
			}
			result = fmt.Sprintf("%s\n|%s|%s|%s|%s|", result, call.ReadableDate, direction, call.Number, call.Duration)
		}
	}

	return result
}
