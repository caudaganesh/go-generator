package structtype

import (
	"go/ast"
	"reflect"
	"strings"

	"github.com/caudaganesh/go-generator/types"
)

type StructType struct{ *ast.StructType }

func GetFromDeclsByName(decls []ast.Decl, name string) *StructType {
	var str *ast.StructType
	for _, decl := range decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		ts, ok := gd.Specs[0].(*ast.TypeSpec)
		if !ok || ts.Name.Name != name {
			continue
		}

		st, ok := ts.Type.(*ast.StructType)
		if ok {
			str = st
		}

	}

	return &StructType{StructType: str}
}

func (s *StructType) GetPropToTag(tag string) PropToTag {
	res := make(PropToTag, len(s.Fields.List))
	for _, f := range s.Fields.List {
		tagValue := f.Tag.Value
		tagValue = strings.Trim(tagValue, "`")
		structTag := reflect.StructTag(tagValue)
		tagValue = structTag.Get(tag)
		res[f.Names[0].Name] = tagValue
	}

	return res
}

func (s *StructType) GetPropToType() PropToType {
	res := make(PropToType, len(s.Fields.List))
	for _, f := range s.Fields.List {
		switch t := f.Type.(type) {
		case *ast.Ident:
			res[f.Names[0].Name] = types.GetPrimitiveType(t.Name, t.Obj)
		}
	}

	return res
}
