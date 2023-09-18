package markdown

const ContactSheetTemplate string = `---
type: person
disabled rules: [ all ]
google:
  resource_name: {{ .ResourceName }}
---

Tags: #contact {{ if .Tags }}{{ .Tags }}{{ end }}

{{ .PersonalData }}{{ if gt (len .Addresses) 0 }}## Address

{{ .Addresses }}
{{ end }}{{ if gt (len .PhoneNumbers) 0 }}## Phone numbers

{{ .PhoneNumbers }}
{{ end }}{{ if gt (len .Email) 0 }}## EMail

{{ .Email }}
{{ end }}{{ if gt (len .Im) 0 }}## IM

{{ .Im }}{{ end }}{{ if gt (len .MessageData) 0 }}## SMS

[[{{.MainLinkName}} Messages]]

{{ end }}{{ if gt (len .CallData.Call) 0 }}## Call log

[[{{.MainLinkName}} Calls]]{{ end }}`
