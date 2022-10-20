package internal

import (
	"text/template"

	"github.com/sascha-andres/sbrdata"
)

func (c *Contact) Handle(pathForFiles, tags string, personalData *template.Template, groups []ContactGroup, addresses, phoneNumbers, emailAddresses, outer *template.Template, sms sbrdata.Messages, calls sbrdata.Calls) {
	var mdData MarkdownData
	mdData.ETag = c.Etag
	mdData.ResourceName = c.ResourceName

	mdData.BuildCalls(calls, c)
	mdData.BuildSms(sms, c)
	mdData.BuildPersonalData(personalData, c)
	mdData.BuildTags(tags, c, groups)
	mdData.BuildAddresses(c, addresses)
	mdData.BuildPhoneNumbers(c, phoneNumbers)
	mdData.BuildEmailAddresses(c, emailAddresses)
	mdData.WriteMarkdown(pathForFiles, outer, c)
}
