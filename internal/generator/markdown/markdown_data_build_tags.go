package markdown

import (
	"strings"

	"github.com/sascha-andres/people2md/internal/types"
)

func (mdData *MarkdownData) BuildTags(tags, tagPrefix string, c *types.Contact, groups []types.ContactGroup) []string {
	result := make([]string, 0)
	for i := range c.Memberships {
		if nil != c.Memberships[i].ContactGroupMembership {
			for _, cg := range groups {
				if cg.ResourceName == c.Memberships[i].ContactGroupMembership.ContactGroupResourceName && strings.Contains(tags, strings.ToLower(cg.Name)) {
					if tagPrefix == "" {
						result = append(result, strings.ToLower(cg.Name))
					} else {
						result = append(result, tagPrefix+strings.ToLower(cg.Name))
					}
				}
			}
		}
	}
	return result
}
