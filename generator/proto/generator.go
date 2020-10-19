package proto

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"

	"github.com/caudaganesh/go-generator/types"
	"github.com/iancoleman/strcase"
)

type Options struct {
	File         string
	TargetStruct string
	PackageName  string
	Name         string
	GoPackage    string
}

func Generate(opt Options) ([]byte, error) {
	src, err := ioutil.ReadFile(opt.File)
	if err != nil {
		return nil, err
	}
	allFields := parseStruct(src, opt.TargetStruct, opt.PackageName)
	properties := getAllProperties(allFields)
	result, err := MakeProto(opt.PackageName, opt.GoPackage, opt.Name, properties)
	return result, err
}

// parseStruct takes in a piece of source code as []byte and get the desired struct fields
func parseStruct(src []byte, structName string, pkgName string) (res []*ast.Field) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", src, 0)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, d := range f.Decls {
		if a, gd := getGenDecl(src, d); a == structName {
			structDecl := gd.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
			res = structDecl.Fields.List
		}
	}

	return
}

func getGenDecl(src []byte, fl interface{}) (string, *ast.GenDecl) {
	gd, ok := fl.(*ast.GenDecl)
	if !ok {
		return "", nil
	}
	t, ok := gd.Specs[0].(*ast.TypeSpec)
	if !ok {
		return "", nil
	}

	return t.Name.Name, gd
}

func getAllProperties(fields []*ast.Field) []string {
	res := []string{}
	for idx, field := range fields {
		decl := fmt.Sprintf("%s %s = %d;", transformFieldTypeToProtoType(field.Type), strcase.ToSnake(field.Names[0].Name), idx+1)
		res = append(res, decl)
	}
	return res
}

func transformFieldTypeToProtoType(fieldType interface{}) string {
	switch ft := fieldType.(type) {
	case *ast.Ident:
		if !types.IsPrimitives(ft.Name) {
			ts, ok := ft.Obj.Decl.(*ast.TypeSpec)
			if ok {
				ident, ok := ts.Type.(*ast.Ident)
				if ok {
					ft.Name = ident.Name
				}
			}
		}

		return TransformTypeToPtype(ft.Name)

	case *ast.SelectorExpr:
		return TransformTypeToPtype(ft.X.(*ast.Ident).Name)
	}

	return ""
}
