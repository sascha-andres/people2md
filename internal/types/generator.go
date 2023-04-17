package types

import (
	"text/template"

	"github.com/sascha-andres/sbrdata"
)

type (
	TemplateIdentifier int

	// DataReferences contains data about what and where
	DataReferences struct {
		Contact      *Contact
		PathForFiles string
		Tags         string
		TagPrefix    string
		Groups       []ContactGroup
		Sms          *sbrdata.Messages
		CallData     *sbrdata.Calls
		Collection   *sbrdata.Collection
	}

	// Templates contains templates used to render
	Templates struct {
		Outer          *template.Template
		Addresses      *template.Template
		PersonalData   *template.Template
		PhoneNumbers   *template.Template
		EmailAddresses *template.Template
		Messages       *template.Template
		Calls          *template.Template
	}

	// Message represents a simple message
	Message struct {
		UnixTimestamp uint64
		Date          string
		Direction     string
		Text          string
	}

	// MessageList is a list of messages to be sorted
	MessageList []Message

	// DataBuilder provides the contract how to construct data
	DataBuilder interface {
		// BuildCalls must not access or rely on fields other than Calls
		BuildCalls(calls sbrdata.CallsData, c *Contact) string
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

// Len of MessageList
func (ml MessageList) Len() int { return len(ml) }

// Swap two messaged
func (ml MessageList) Swap(i, j int) { ml[i], ml[j] = ml[j], ml[i] }

// Less determines whether i < j for a message (unix timestamp used)
func (ml MessageList) Less(i, j int) bool { return ml[i].UnixTimestamp < ml[j].UnixTimestamp }
