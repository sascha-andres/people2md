package internal

type (
	MarkdownData struct {
		ETag         string
		ResourceName string
		PersonalData string
		Addresses    string
		Im           string
		PhoneNumbers string
		Email        string
		Tags         string
	}

	ContactGroup struct {
		Etag          string
		FormattedName string
		GroupType string
		MetaData  ContactGroupMetaData
		Name      string
		ResourceName  string
	}

	ContactGroupMetaData struct {
		UpdateTime string
	}

	Contact struct {
		Etag           string
		Memberships    []Membership
		Names          []Name
		PhoneNumbers   []PhoneNumber
		ResourceName   string
		EmailAddresses []EmailAddress
		Organizations  []Organization
		ImClients      []ImClient
		Birthdays      []Birthday
		Addresses      []Address
	}

	Address struct {
		MetaData        *MetaData
		FormattedType   string
		FormattedValue  string
		Type            string
		City            string
		Country         string
		ExtendedAddress string
		PostalCode      string
		StreetAddress   string
	}

	Birthday struct {
		MetaData *MetaData
		Date     *Date
		Text     *string
	}

	Date struct {
		Day   uint
		Month uint
		Year  uint
	}

	ImClient struct {
		FormattedProtocol string
		MetaData          *MetaData
		Protocol          string
		Username          string
	}

	Organization struct {
		FormattedType string
		MetaData      *MetaData
		Name          string
		Type          string
		Title         string
	}

	EmailAddress struct {
		FormattedType string
		MetaData      *MetaData
		Type          string
		Value         string
	}

	PhoneNumber struct {
		CanonicalForm string
		FormattedType string
		MetaData      *MetaData
		Value         string
		Type          string
	}

	Name struct {
		DisplayName          string
		DisplayNameLastFirst string
		FamilyName           string
		GivenName            string
		MetaData             *MetaData
		MiddleName           string
		UnstructuredName     string
	}

	Membership struct {
		ContactGroupMembership *ContactGroupMembership
		MetaData               *MetaData
	}

	ContactGroupMembership struct {
		ContactGroupId           string
		ContactGroupResourceName string
	}

	MetaData struct {
		Primary bool
		Source  *Source
	}

	Source struct {
		Id   string
		Type string
	}
)
