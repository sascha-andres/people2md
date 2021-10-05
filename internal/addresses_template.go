package internal

const AddressesTemplate string = `{{ range . }}{{ if .FormattedType }}{{ .FormattedType }}: {{ end }}{{ .FormattedValue }}
{{ end }}`