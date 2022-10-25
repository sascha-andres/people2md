package markdown

const MarkDownTemplate string = `---
type: person
disabled rules: [ all ]
google:
  resource_name: {{ .ResourceName }}
  etag: {{ .ETag }}
---

Tags: #contact {{ if .Tags }}{{ .Tags }}{{ end }}

{{ .PersonalData }}{{ if gt (len .Addresses) 0 }}## Address

{{ .Addresses }}
{{ end }}{{ if gt (len .PhoneNumbers) 0 }}## Phone numbers

{{ .PhoneNumbers }}
{{ end }}{{ if gt (len .Email) 0 }}## EMail

{{ .Email }}
{{ end }}{{ if gt (len .Im) 0 }}## IM

{{ .Im }}{{ end }}{{ if gt (len .Sms) 0 }}## SMS

{{ .Sms }}

{{ end }}{{ if gt (len .Calls) 0 }}## Call log

{{ .Calls }}{{ end }}`
