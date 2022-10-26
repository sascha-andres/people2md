package markdown

import (
	"fmt"
	"github.com/sascha-andres/people2md/internal/types"
	"strings"
)

func (mdData *MarkdownData) BuildSms(ml types.MessageList) string {
	result := ""

	for _, message := range ml {
		if result == "" {
			result = `|Date|Direction|Text|
|---|---|---|`
		}
		result = fmt.Sprintf("%s\n|%s|%s|%s|", result, message.Date, message.Direction, sanitizeBody(message.Text))
	}

	return result
}

func sanitizeBody(body string) string {
	result := strings.Replace(body, "<", "&lt;", -1)
	result = strings.Replace(result, ">", "&gt;", -1)
	result = strings.Replace(result, "\n", "<br />", -1)
	return result
}
