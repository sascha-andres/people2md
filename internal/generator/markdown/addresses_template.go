package markdown

const AddressesTemplate string = `{{ range . }}{{ if .FormattedType }}{{ .FormattedType }}: {{ end }}{{ .FormattedValue }}
{{ end }}`
