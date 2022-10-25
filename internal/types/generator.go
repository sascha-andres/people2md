package types

import (
	"github.com/sascha-andres/sbrdata"
	"text/template"
)

type (
	TemplateIdentifier int

	Templates struct {
		Outer          *template.Template
		Addresses      *template.Template
		PersonalData   *template.Template
		PhoneNumbers   *template.Template
		EmailAddresses *template.Template
	}

	DataBuilder interface {
		BuildCalls(calls sbrdata.Calls, c *Contact) string
		BuildSms(calls sbrdata.Messages, c *Contact) string
		BuildPersonalData(personalData *template.Template, c *Contact) string
		BuildTags(tags string, c *Contact, groups []ContactGroup) string
		BuildAddresses(c *Contact, addresses *template.Template) string
		BuildPhoneNumbers(c *Contact, phoneNumbers *template.Template) string
		BuildEmailAddresses(c *Contact, emailAddresses *template.Template) string
		SetETag(etag string)
		SetResourceName(rn string)
		GetTemplate(id TemplateIdentifier) *template.Template
		GetTemplateData(id TemplateIdentifier) []byte
	}
)
