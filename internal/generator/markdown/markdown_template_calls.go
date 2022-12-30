package markdown

const MarkDownTemplateCalls string = `---
type: calls
disabled rules: [ all ]
google:
  resource_name: {{ .ResourceName }}
  etag: {{ .ETag }}
---

Tags: #contact {{ if .Tags }}{{ .Tags }}{{ end }}

[[{{.MainLinkName}}]]

# Call log

{{ .Calls }}`
