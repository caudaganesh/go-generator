package interfacegen

import (
	"bytes"
	"html"
	"html/template"
	"log"

	"golang.org/x/tools/imports"
)

type templateParam struct {
	Name        string
	PackageName string
	Methods     []string
	Comment     string
}

func MakeInterface(pkgName, name, comment string, methods []string) ([]byte, error) {
	baseTemplate := `
	package {{.PackageName}}
	
	// {{.Comment}}
	type {{.Name}} interface {
		{{range $_, $method := .Methods}}
		{{$method}}
		{{- end}}
	}
	`

	tmpl := template.Must(template.New("interface").Parse(baseTemplate))
	var buf bytes.Buffer
	tp := templateParam{
		PackageName: pkgName,
		Name:        name,
		Methods:     methods,
		Comment:     comment,
	}
	err := tmpl.Execute(&buf, tp)
	if err != nil {
		log.Fatal(err)
	}
	res := buf.String()
	res = html.UnescapeString(res)
	return format(res)
}

func format(code string) ([]byte, error) {
	opts := &imports.Options{
		TabIndent: true,
		TabWidth:  4,
		Fragment:  true,
		Comments:  true,
	}
	return imports.Process("", []byte(code), opts)
}
