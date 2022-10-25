package markdown

const EmailsTemplate string = `{{ range . }}{{ .FormattedType }}: {{ .Value }}
{{ end }}`
