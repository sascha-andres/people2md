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

|Date|Direction|Number|Duration (s)|
|---|---|---|---|
{{ range .CallData.Call }}|{{ .ReadableDate }}|{{ if eq .Type "2" }}outgoing{{ else }}incoming{{ end }}|{{ .Number }}|{{ .Duration }}|
{{ end }}
`
