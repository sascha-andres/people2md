package markdown

import (
	"fmt"
	"strings"

	"github.com/sascha-andres/people2md/internal/types"
	"github.com/sascha-andres/sbrdata"
)

func (mdData *MarkdownData) BuildCalls(calls sbrdata.CallsData, c *types.Contact) string {
	result := ""

	for _, call := range calls.GetCalls() {
		include := false
		for _, name := range c.Names {
			include = name.DisplayName == call.GetContactName()
			if include {
				break
			}
		}
		if !include {
			for _, org := range c.Organizations {
				include = org.Name == call.GetContactName()
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
		if !include && call.GetNumber() != "" {
			num := call.GetNumber()
			if strings.HasPrefix(num, "0") {
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
			result = fmt.Sprintf("%s\n|%s|%s|%s|%s|", result, call.GetReadableDate(), direction, call.GetNumber(), call.GetDuration())
		}
	}

	return result
}
