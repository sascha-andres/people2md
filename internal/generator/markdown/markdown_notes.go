package markdown

const NotesSheetTemplate string = `---
type: person
disabled rules: [ all ]
google:
  resource_name: {{ .ResourceName }}
tags:
  - contact{{ if gt (len .Tags) 0 }}{{ range .Tags }}
  - {{ . }}{{ end }}{{ end }}
---

return to [[{{.MainLinkName}}]]`
