package main

const entryTmpl = `
entity "{{ .Name }}" {
{{- if .Comment.Valid }}
  {{ .Comment.String }}
  ..
{{- end }}
{{- range .Columns }}
  {{- if .IsPrimaryKey }}
  + {{ .Name }}: type({{ .DataType }}) [PK]{{if .IsForeignKey }}[FK]{{end}}{{- if .Comment.Valid }} : {{ .Comment.String }}{{- end }}
  {{- end }}
{{- end }}
  --
{{- range .Columns }}
  {{- if not .IsPrimaryKey }}
  {{if .NotNull}}*{{end}}{{ .Name }}: type({{ .DataType }}) {{if .IsForeignKey}}[FK]{{end}} {{- if .Comment.Valid }} : {{ .Comment.String }}{{- end }}
  {{- end }}
{{- end }}
}
`

const relationTmpl = `
{{ if .IsOneToOne }} {{ .SourceTableName }} ||-|| {{ .TargetTableName }}{{else}} {{ .SourceTableName }} }-- {{ .TargetTableName }}{{end}}
`

const enumTmpl = `
enum "{{ .Name }}" {
{{- range .Values }}
	{{ .Value }}: {{ .Order }}	
{{- end }}
}
`
