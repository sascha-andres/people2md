package markdown

const MarkDownTemplateMessages string = `---
type: messages
disabled rules: [ all ]
google:
  resource_name: {{ .ResourceName }}
  etag: {{ .ETag }}
---

Tags: #contact {{ if .Tags }}{{ .Tags }}{{ end }}

[[{{.MainLinkName}}]]

# Message log

|Date|Direction|Text|
|---|---|---|
{{ range .MessageData }}|{{ .Date }}|{{ if eq .Direction "2" }}outgoing{{ else }}incoming{{ end }}|{{ .Text }}|
{{ end }}`
