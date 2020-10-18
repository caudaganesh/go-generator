package interfacegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type Method struct {
	Code string
	Docs []string
}

type Options struct {
	File         string
	TargetStruct string
	PackageName  string
	Name         string
	Comment      string
}

func Generate(opt Options) ([]byte, error) {
	src, err := ioutil.ReadFile(opt.File)
	if err != nil {
		return nil, err
	}

	allMethods := getAllMethods(opt.PackageName, src, opt.TargetStruct)
	result, err := MakeInterface(opt.PackageName, opt.Name, opt.Comment, allMethods)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getAllMethods(
	pkgName string,
	src []byte,
	structType string) []string {
	allMethods := []string{}
	methods := parseStruct(src, structType, pkgName)
	for _, m := range methods {
		allMethods = append(allMethods, m.lines()...)
	}

	return allMethods
}

func parseStruct(src []byte, structName string, pkgName string) (methods []Method) {
	fs := token.NewFileSet()
	a, err := parser.ParseFile(fs, "", src, parser.ParseComments)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, d := range a.Decls {
		a, fd := getNameAndFuncDecl(src, d)
		if a != structName || !fd.Name.IsExported() {
			continue
		}

		params := formatFieldList(src, fd.Type.Params, pkgName)
		rets := formatFieldList(src, fd.Type.Results, pkgName)
		method := fmt.Sprintf("%s(%s) (%s)", fd.Name.String(), strings.Join(params, ", "), strings.Join(rets, ", "))
		docs := getDocs(fd, src)
		methods = append(methods, Method{Code: method, Docs: docs})
	}

	return
}

func getNameAndFuncDecl(src []byte, fl interface{}) (string, *ast.FuncDecl) {
	fd, ok := fl.(*ast.FuncDecl)
	if !ok {
		return "", nil
	}

	t, err := getReceiverType(fd)
	if err != nil {
		return "", nil
	}

	st := string(src[t.Pos()-1 : t.End()-1])
	if len(st) > 0 && st[0] == '*' {
		st = st[1:]
	}

	return st, fd
}

func getReceiverType(fd *ast.FuncDecl) (ast.Expr, error) {
	if fd.Recv == nil {
		return nil, fmt.Errorf("fd is not a method, it is a function")
	}

	return fd.Recv.List[0].Type, nil
}

func formatFieldList(src []byte, fl *ast.FieldList, pkgName string) []string {
	if fl == nil {
		return nil
	}

	var parts []string
	for _, l := range fl.List {
		names := make([]string, len(l.Names))
		for i, n := range l.Names {
			names[i] = n.Name
		}

		t := string(src[l.Type.Pos()-1 : l.Type.End()-1])
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

func getDocs(fd *ast.FuncDecl, src []byte) []string {
	if fd.Doc == nil {
		return nil
	}

	var docs []string
	for _, d := range fd.Doc.List {
		docs = append(docs, string(src[d.Pos()-1:d.End()-1]))
	}

	return docs
}

func (m *Method) lines() []string {
	var lines []string
	lines = append(lines, m.Docs...)
	lines = append(lines, m.Code)
	return lines
}
