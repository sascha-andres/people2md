package markdown

const PhoneNumbersTemplate string = `{{ range . }}{{ .FormattedType }}: {{ if .CanonicalForm }}{{ .CanonicalForm }}{{ else }}{{ if .Value }}{{ .Value }}{{ end }}{{ end }}
{{ end }}`