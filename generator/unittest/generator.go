package unittest

import (
	"bytes"
	"fmt"
	"go/token"
	"io"
	"text/template"

	"golang.org/x/tools/imports"
)

type Options struct {
	Comment        string
	Template       string
	HeaderTemplate string
	PackageName    string
	PkgPath        string

	Funcs []*Func
}

type TemplateData struct {
	Comment       string
	Func          *Func
	CalledMethods []CalledMethod
}

type CalledMethod struct {
	FieldName     string
	MethodName    string
	Args, Results []string
}

//Generator is used to generate a test stub for function Func
type Generator struct {
	opt            Options
	buf            *bytes.Buffer
	headerTemplate *template.Template
	testTemplate   *template.Template
}

//NewGenerator returns a pointer to Generator
func NewGenerator(opt Options) *Generator {
	fs := token.NewFileSet()
	return &Generator{
		buf:            bytes.NewBuffer([]byte{}),
		opt:            opt,
		headerTemplate: template.Must(template.New("header").Funcs(templateHelpers(fs)).Parse(opt.HeaderTemplate)),
		testTemplate:   template.Must(template.New("test").Funcs(templateHelpers(fs)).Parse(opt.Template)),
	}
}

func (g *Generator) Write(w io.Writer) error {
	if len(g.opt.Funcs) == 0 {
		return nil
	}

	if g.buf.Len() == 0 {
		if err := g.WriteHeader(g.buf); err != nil {
			return err
		}
	}

	if err := g.WriteTests(g.buf); err != nil {
		return err
	}

	formattedSource, err := imports.Process("", g.buf.Bytes(), nil)
	if err != nil {
		return err
	}

	if _, err = w.Write(formattedSource); err != nil {
		return err
	}

	return nil
}

func (g *Generator) Source() string {
	return g.buf.String()
}

//WriteHeader writes a package name and import specs
func (g *Generator) WriteHeader(w io.Writer) error {
	return g.headerTemplate.Execute(w, struct {
		Package string
		Comment string
	}{
		Package: g.opt.PackageName,
		Comment: g.opt.Comment,
	})
}

//WriteTests writes test stubs for every function that don't have test yet
func (g *Generator) WriteTests(w io.Writer) error {
	for _, f := range g.opt.Funcs {

		templateData := TemplateData{
			Comment: g.opt.Comment,
			Func:    f,
		}
		for _, expr := range f.FindCalledMethods() {
			args, result, err := expr.GetArgsReturns()
			if err != nil {
				// do we need handle error?
				continue
			}
			templateData.CalledMethods = append(templateData.CalledMethods,
				CalledMethod{
					FieldName:  expr.FieldName(),
					MethodName: expr.Name(),
					Args:       args,
					Results:    result,
				})
		}

		err := g.testTemplate.Execute(w, templateData)

		if err != nil {
			return fmt.Errorf("failed to write test: %v", err)
		}
	}

	return nil
}
