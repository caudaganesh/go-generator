package structtype

import "go/ast"

func GetFromDeclsByName(decls []ast.Decl, name string) *ast.StructType {
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

	return str
}
