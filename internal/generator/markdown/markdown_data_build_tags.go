package markdown

import (
	"fmt"
	"strings"

	"github.com/sascha-andres/people2md/internal/types"
)

func (mdData *MarkdownData) BuildTags(tags, tagPrefix string, c *types.Contact, groups []types.ContactGroup) string {
	result := ""
	for i := range c.Memberships {
		if nil != c.Memberships[i].ContactGroupMembership {
			for _, cg := range groups {
				if cg.ResourceName == c.Memberships[i].ContactGroupMembership.ContactGroupResourceName && strings.Contains(tags, strings.ToLower(cg.Name)) {
					if result == "" {
						result = fmt.Sprintf("#%s%s", tagPrefix, strings.ToLower(cg.Name))
					} else {
						result = fmt.Sprintf("%s #%s%s", result, tagPrefix, strings.ToLower(cg.Name))
					}
				}
			}
		}
	}
	return result
}
