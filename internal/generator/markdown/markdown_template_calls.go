package markdown

const MarkDownTemplateCalls string = `---
type: calls
disabled rules: [ all ]
date created: {{ .DateCreated }}
date modified: {{ .DateModified }}
google:
  resource_name: {{ .ResourceName }}
---

Tags: #contact {{ if .Tags }}{{ .Tags }}{{ end }}

[[{{.MainLinkName}}]]

# Call log

|Date|Direction|Number|Duration (s)|
|---|---|---|---|
{{ range .CallData.Call }}|{{ .ReadableDate }}|{{ if eq .Type "2" }}outgoing{{ else }}incoming{{ end }}|{{ .Number }}|{{ .Duration }}|
{{ end }}`
