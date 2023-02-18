package diff

const (
	fieldsKey = "Fields"

	defTmpl = `
	{{- range .Fields -}}
	{{.name}} changed: before:{{.before}} after:{{.after}}
	{{end}}`
)
