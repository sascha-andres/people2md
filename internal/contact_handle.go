package internal

import "text/template"

func (c *Contact) Handle(tags string, personalData *template.Template, groups []ContactGroup, addresses *template.Template, phoneNumbers *template.Template, emailAddresses *template.Template, outer *template.Template) {
	var mdData MarkdownData
	mdData.ETag = c.Etag
	mdData.ResourceName = c.ResourceName

	mdData.BuildPersonalData(personalData, c)
	mdData.BuildTags(tags, c, groups)
	mdData.BuildAddresses(c, addresses)
	mdData.BuildPhoneNumbers(c, phoneNumbers)
	mdData.BuildEmailAddresses(c, emailAddresses)
	mdData.WriteMarkdown(outer, c)
}
