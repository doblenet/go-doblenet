package cliapp

import (
	"io"
	"text/tabwriter"
	"text/template"
)

const (
	versionTemplate = `{{.Name}} v{{.Version}}{{if .GitTag}}+git{{.GitTag}}{{end}}
`
	attribTemplate = `{{if len .}}
ATTRIBUTES:
{{range $key,$value := .}}	{{$key}}:	{{$value}}
{{end}}{{end}}`
	infoTemplate = `{{if len .Authors}}
AUTHOR(S):
	{{range .Authors}}{{.}}
	{{end}}{{end}}{{if .Copyright}}
COPYRIGHT:
	{{.Copyright}}
{{if .Email}}
	bugs to {{.Email}}{{end}}
{{end}}
`
)


func printVersionInfo(a *App, out io.Writer) {
	renderText(out, "version", versionTemplate, a.info)
}

func printExtendedInfo(a *App, out io.Writer) {
	
	// Initialize a tabWriter connected to "stdout" ...
	w := tabwriter.NewWriter(out, 1, 8, 2, ' ', 0)
	
	// Version info
	renderText(w, "version", versionTemplate, a.info)
	
	// Attributes
	renderText(w, "attribs", attribTemplate, a.metadata)

	// Authors & copyright
	ci := struct{Authors []Author ; Copyright string; Email string}{a.authors, a.info.Copyright, a.info.Email}
	renderText(w, "copyright", infoTemplate, ci)

	w.Flush()
}


func renderText(w io.Writer, title string, tmpl string, data interface{}) {
// 	funcMap := template.FuncMap{
// 		"join": strings.Join,
// 	}

	t := template.Must(template.New(title).Parse(tmpl))
// 	t := template.Must(template.New(title).Funcs(funcMap).Parse(tmpl))
	
	if err := t.Execute(w, data); nil != err {
		// If the writer is closed, t.Execute will fail, and there's nothing
		// we can do to recover.
		return
	}
	return
}
