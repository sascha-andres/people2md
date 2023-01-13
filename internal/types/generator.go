package types

import (
	"text/template"

	"github.com/sascha-andres/sbrdata"
)

type (
	TemplateIdentifier int

	Templates struct {
		Outer          *template.Template
		Addresses      *template.Template
		PersonalData   *template.Template
		PhoneNumbers   *template.Template
		EmailAddresses *template.Template
		Messages       *template.Template
		Calls          *template.Template
	}

	Message struct {
		UnixTimestamp uint64
		Date          string
		Direction     string
		Text          string
	}

	MessageList []Message

	DataBuilder interface {
		BuildCalls(calls sbrdata.Calls, c *Contact) string
		BuildMessages(messages MessageList) string
		BuildPersonalData(personalData *template.Template, c *Contact) string
		BuildTags(tags, tagPrefix string, c *Contact, groups []ContactGroup) string
		BuildAddresses(c *Contact, addresses *template.Template) string
		BuildPhoneNumbers(c *Contact, phoneNumbers *template.Template) string
		BuildEmailAddresses(c *Contact, emailAddresses *template.Template) string
		SetETag(etag string)
		SetResourceName(rn string)
		GetTemplate(id TemplateIdentifier) *template.Template
		GetTemplateData(id TemplateIdentifier) []byte
	}
)

func (ml MessageList) Len() int           { return len(ml) }
func (ml MessageList) Swap(i, j int)      { ml[i], ml[j] = ml[j], ml[i] }
func (ml MessageList) Less(i, j int) bool { return ml[i].UnixTimestamp < ml[j].UnixTimestamp }
