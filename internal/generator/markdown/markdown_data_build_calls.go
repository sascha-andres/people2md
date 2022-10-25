package markdown

import (
	"fmt"
	"github.com/sascha-andres/people2md/internal/types"
	"github.com/sascha-andres/sbrdata"
)

func (mdData *MarkdownData) BuildCalls(calls sbrdata.Calls, c *types.Contact) string {
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
