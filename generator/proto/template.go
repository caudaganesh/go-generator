package proto

import (
	"bytes"
	"html"
	"log"
	"text/template"
)

type templateParam struct {
	Name        string
	PackageName string
	GoPackage   string
	Properties  []string
}

// MakeProto takes in all of the fields and parse it to proto
func MakeProto(pkgName, goPackage, name string, properties []string) ([]byte, error) {
	baseTemplate := `
syntax="proto3";

package {{.PackageName}};

{{if .GoPackage}}option go_package = {{.GoPackage}};{{end}}

message {{.Name}} {
	{{range $_, $prop := .Properties}}
	{{$prop}}
	{{- end}}
}`

	tmpl := template.Must(template.New("proto").Parse(baseTemplate))
	var buf bytes.Buffer
	tp := templateParam{
		PackageName: pkgName,
		Name:        name,
		Properties:  properties,
		GoPackage:   goPackage,
	}

	err := tmpl.Execute(&buf, tp)
	if err != nil {
		log.Fatal(err)
	}

	res := buf.String()
	res = html.UnescapeString(res)
	return []byte(res), nil
}
