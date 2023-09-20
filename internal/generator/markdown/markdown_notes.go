package markdown

const NotesSheetTemplate string = `---
type: person
disabled rules: [ all ]
google:
  resource_name: {{ .ResourceName }}
---

Tags: #contact {{ if .Tags }}{{ .Tags }}{{ end }}

return to [[{{.MainLinkName}}]]`
