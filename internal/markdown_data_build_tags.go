package internal

import (
	"fmt"
	"strings"
)

func (mdData *MarkdownData) BuildTags(tags string, c *Contact, groups []ContactGroup) {
	for i := range c.Memberships {
		if nil != c.Memberships[i].ContactGroupMembership {
			for _, cg := range groups {
				if cg.ResourceName == c.Memberships[i].ContactGroupMembership.ContactGroupResourceName && strings.Contains(tags, strings.ToLower(cg.Name)) {
					if mdData.Tags == "" {
						mdData.Tags = fmt.Sprintf("#%s", strings.ToLower(cg.Name))
					} else {
						mdData.Tags = fmt.Sprintf("%s #%s", mdData.Tags, strings.ToLower(cg.Name))
					}
				}
			}
		}
	}
}
