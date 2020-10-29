package unittest

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"strings"
	"text/template"

	"github.com/caudaganesh/go-generator/pkgloader"
)

type matchFunc func(*ast.FuncDecl) bool

func GetFunctionsFromFile(file *ast.File, pkgPath string) []*Func {
	funcs := findFunctions(file, func(fd *ast.FuncDecl) bool {
		return true
	}, pkgPath)

	return funcs
}

//findMissingTests filters funcs slice and returns only those functions that don't have tests yet
func findMissingTests(file *ast.File, funcs []*Func) []*Func {
	tests := findFunctions(file, func(fd *ast.FuncDecl) bool {
		for _, sourceFunc := range funcs {
			f := NewFunc(fd, nil, file, nil, "") //TODO: fix this later
			if f.ReceiverType() == nil && f.Name() == sourceFunc.TestName() {
				return true
			}
		}
		return false
	}, "")

	dontHaveTests := []*Func{}
	for _, f := range funcs {
		testIsFound := false
		for _, test := range tests {
			if test.Name() == f.TestName() {
				testIsFound = true
				break
			}
		}
		if !testIsFound {
			dontHaveTests = append(dontHaveTests, f)
		}
	}

	return dontHaveTests
}

//findFunctions finds all matching function declarations
func findFunctions(file *ast.File, match matchFunc, pkgPath string) []*Func {
	var funcs []*Func

	// required because it's possible that the struct and the method are in different files
	pkgDecls := getPackageDecls(pkgPath)

	for _, decl := range file.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || !match(fd) {
			continue
		}

		// will handle non-method / functions without receiver
		if fd.Recv == nil {
			funcs = append(funcs, NewFunc(fd, nil, file, pkgDecls, pkgPath))
			continue
		}

		stExp, ok := fd.Recv.List[0].Type.(*ast.StarExpr)
		if !ok {
			continue
		}

		id, ok := stExp.X.(*ast.Ident)
		if !ok {
			continue
		}

		strType := getStructType(pkgDecls, id.Name)
		if fd.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Obj == nil {
			fd.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Obj = generateTypeObject(id, strType)
		}

		funcs = append(funcs, NewFunc(fd, strType, file, pkgDecls, pkgPath))
	}

	return funcs
}

func getPackageDecls(pkgPath string) []ast.Decl {
	pkg, err := pkgloader.Load(pkgPath)
	if err != nil {
		log.Fatal(err)
	}

	var pkgDecls []ast.Decl
	for _, f := range pkg.Syntax {
		pkgDecls = append(pkgDecls, f.Decls...)
	}

	return pkgDecls
}

func getStructType(decls []ast.Decl, name string) *ast.StructType {
	recStructName := name
	var strType *ast.StructType
	for _, d := range decls {
		gd, ok := d.(*ast.GenDecl)
		if !ok {
			continue
		}

		typ, ok := gd.Specs[0].(*ast.TypeSpec)
		if !ok {
			continue
		}

		strType, ok = typ.Type.(*ast.StructType)
		if ok && recStructName == typ.Name.Name {
			break
		}
	}

	return strType
}

func generateTypeObject(id *ast.Ident, strType *ast.StructType) *ast.Object {
	return &ast.Object{
		Kind: ast.Typ,
		Name: id.Name,
		Decl: &ast.TypeSpec{
			Doc:     &ast.CommentGroup{},
			Name:    id,
			Assign:  0,
			Type:    strType,
			Comment: &ast.CommentGroup{},
		},
		Data: nil,
		Type: nil,
	}
}

//nodeToString returns a string representation of an AST node
//as it has in the original source code
func nodeToString(fs *token.FileSet, n ast.Node) string {
	b := bytes.NewBuffer([]byte{})
	printer.Fprint(b, fs, n)
	return b.String()
}

//templateHelpers return FuncMap of template helpers to use within a template
func templateHelpers(fs *token.FileSet) template.FuncMap {
	return template.FuncMap{
		"ast": func(n ast.Node) string {
			return nodeToString(fs, n)
		},
		"join": strings.Join,
		"params": func(f *Func) []string {
			return f.Params(fs)
		},
		"fields": func(f *Func) []string {
			return f.Fields(fs)
		},
		"results": func(f *Func) []string {
			return f.Results(fs)
		},
		"receiver": func(f *Func) string {
			if f.ReceiverType() == nil {
				return ""
			}

			return strings.Replace(nodeToString(fs, f.ReceiverType()), "*", "", -1) + "."
		},
		"want": func(s string) string { return strings.Replace(s, "got", "want", 1) },
	}
}
