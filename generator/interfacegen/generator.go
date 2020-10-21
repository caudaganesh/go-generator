package interfacegen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/caudaganesh/go-generator/pkgloader"
)

type Method struct {
	Code string
	Docs []string
}

type Options struct {
	Package      string
	File         string
	TargetStruct string
	PackageName  string
	Name         string
	Comment      string
}

func Generate(opt Options) ([]byte, error) {
	allMethods := opt.getAllMethods(opt.PackageName, opt.TargetStruct)
	result, err := MakeInterface(opt.PackageName, opt.Name, opt.Comment, allMethods)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *Options) GetFileSetAndDecls() (*token.FileSet, []ast.Decl) {
	if o.File != "" {
		return o.getFSetAndDeclsByFile()
	}

	if o.Package != "" {
		return o.getFSetAndDeclsByPackage()
	}

	return nil, nil
}

func (o *Options) getFSetAndDeclsByFile() (*token.FileSet, []ast.Decl) {
	src, err := ioutil.ReadFile(o.File)
	if err != nil {
		log.Fatal(err)
	}

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", src, parser.ParseComments)
	if err != nil {
		log.Fatal(err.Error())
	}

	return fs, f.Decls
}

func (o *Options) getFSetAndDeclsByPackage() (*token.FileSet, []ast.Decl) {
	p, err := pkgloader.Load(o.Package)
	if err != nil {
		log.Fatal(err)
	}

	var decls []ast.Decl
	for _, f := range p.Syntax {
		decls = append(decls, f.Decls...)
	}

	return p.Fset, decls
}

func (o *Options) getAllMethods(
	pkgName string,
	structType string) []string {
	allMethods := []string{}
	fs, decls := o.GetFileSetAndDecls()
	methods := parseDecls(fs, decls, structType, pkgName)
	for _, m := range methods {
		allMethods = append(allMethods, m.lines()...)
	}

	return allMethods
}

func getDecls(files []*ast.File) []ast.Decl {
	var decls []ast.Decl

	for _, file := range files {
		decls = append(decls, file.Decls...)
	}

	return decls
}

func parseDecls(fs *token.FileSet, decls []ast.Decl, structName string, pkgName string) (methods []Method) {

	for _, decl := range decls {
		st, fd := getNameAndFuncDecl(decl)
		if st != structName || !fd.Name.IsExported() {
			continue
		}
		params := formatFieldList(fs, fd.Type.Params, pkgName)
		rets := formatFieldList(fs, fd.Type.Results, pkgName)
		method := fmt.Sprintf("%s(%s) (%s)", fd.Name.String(), strings.Join(params, ", "), strings.Join(rets, ", "))
		docs := getDocs(fs, fd)
		methods = append(methods, Method{Code: method, Docs: docs})
	}

	return
}

func getNameAndFuncDecl(fl interface{}) (string, *ast.FuncDecl) {
	fd, ok := fl.(*ast.FuncDecl)
	if !ok {
		return "", nil
	}

	if fd.Recv == nil {
		return "", fd
	}

	st := fd.Recv.
		List[0].
		Type.(*ast.StarExpr).
		X.(*ast.Ident).Name

	return st, fd
}

func formatFieldList(fs *token.FileSet, fl *ast.FieldList, pkgName string) []string {
	if fl == nil {
		return nil
	}

	var parts []string
	for _, l := range fl.List {
		names := make([]string, len(l.Names))
		for i, n := range l.Names {
			names[i] = n.Name
		}

		var tb bytes.Buffer
		printer.Fprint(&tb, fs, l.Type)
		t := tb.String()
		regexString := fmt.Sprintf(`(\*|\(|\s|^)%s\.`, regexp.QuoteMeta(pkgName))
		t = regexp.MustCompile(regexString).ReplaceAllString(t, "$1")

		if len(names) > 0 {
			typeSharingArgs := strings.Join(names, ", ")
			parts = append(parts, fmt.Sprintf("%s %s", typeSharingArgs, t))
		} else {
			parts = append(parts, t)
		}
	}
	return parts
}

func getDocs(fs *token.FileSet, fd *ast.FuncDecl) []string {
	if fd.Doc == nil {
		return nil
	}

	var docs []string
	for _, d := range fd.Doc.List {
		docs = append(docs, d.Text)
	}

	return docs
}

func (m *Method) lines() []string {
	var lines []string
	lines = append(lines, m.Docs...)
	lines = append(lines, m.Code)
	return lines
}
