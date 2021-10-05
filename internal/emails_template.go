package internal

const EmailsTemplate string = `{{ range . }}{{ .FormattedType }}: {{ .Value }}
{{ end }}`
