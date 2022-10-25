package markdown

const PersonalDataTemplate string = `{{ if gt (len .Names) 0 }}# {{ (index .Names 0).DisplayName }}
{{ if gt (len .Organizations) 0 }}
{{ range .Organizations }}__{{ .Name }}__{{ if .Title }} ({{ .Title }}){{end}}{{end}}
{{ end }}
Notes: [[{{ (index .Names 0).DisplayName }} Notes|Notes]]

{{ if gt (len .Birthdays) 0 }}{{ range .Birthdays }}Birthday: {{ if .Date }}{{ .Date.Year }}-{{ .Date.Month }}-{{ .Date.Day }}{{ else }}{{ if .Text }}{{ .Text }}{{end}}{{ end }} 
{{end}}
{{ end }}{{ else }}{{ if gt (len .Organizations) 0 }}# {{ (index .Organizations 0).Name }}
Notes: [[{{ (index .Organizations 0).Name }} Notes|Notes]]

{{ end }}{{ end }}`
