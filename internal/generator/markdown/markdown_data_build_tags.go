package markdown

import (
	"fmt"
	"github.com/sascha-andres/people2md/internal/types"
	"strings"
)

func (mdData *MarkdownData) BuildTags(tags string, c *types.Contact, groups []types.ContactGroup) string {
	result := ""
	for i := range c.Memberships {
		if nil != c.Memberships[i].ContactGroupMembership {
			for _, cg := range groups {
				if cg.ResourceName == c.Memberships[i].ContactGroupMembership.ContactGroupResourceName && strings.Contains(tags, strings.ToLower(cg.Name)) {
					if result == "" {
						result = fmt.Sprintf("#%s", strings.ToLower(cg.Name))
					} else {
						result = fmt.Sprintf("%s #%s", result, strings.ToLower(cg.Name))
					}
				}
			}
		}
	}
	return result
}
