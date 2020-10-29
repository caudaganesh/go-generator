package unittest

import (
	"go/ast"
	"go/types"
	"log"
	"strings"

	"github.com/caudaganesh/go-generator/pkgloader"
)

type ReceiverField struct {
	CalledMethods []*CalledExpr
	Field         *ast.Field
	PkgPath       string
	Name          string
	StructName    string
	Mockable      bool
	IsCalled      bool
}

func NewReceiverField(
	name,
	pkgPath string,
	fd *ast.Field,
	calledMethods []*CalledExpr,
) ReceiverField {
	return ReceiverField{
		Field:         fd,
		PkgPath:       pkgPath,
		Name:          name,
		CalledMethods: calledMethods,
	}
}

func (r *ReceiverField) Set() {
	r.IsCalled = r.IsNameCalled()
	r.Mockable = r.IsFieldMockable()
	switch typ := r.Field.Type.(type) {
	case *ast.SelectorExpr:
		r.StructName = typ.Sel.Name

	case *ast.Ident:
		r.StructName = typ.Name

	case *ast.ArrayType:
		r.SetForArray(typ.Elt)
	}
}

func (r *ReceiverField) SetForArray(typ ast.Expr) {
	switch e := typ.(type) {
	case *ast.SelectorExpr:
		r.StructName = e.Sel.Name

	case *ast.Ident:
		r.StructName = e.Name
	}
}

func (r *ReceiverField) IsNameCalled() bool {
	for _, expr := range r.CalledMethods {
		if expr.FieldName() == r.Name {
			return true
		}
	}

	return false
}

func (r *ReceiverField) IsFieldMockable() bool {
	switch typ := r.Field.Type.(type) {
	case *ast.SelectorExpr:
		return r.IsSelectorExprMockable(typ)
	case *ast.Ident:
		return r.IsIdentMockable(typ)
	}

	return false
}

func (r *ReceiverField) IsSelectorExprMockable(expr *ast.SelectorExpr) bool {
	pkg, err := pkgloader.Load(r.PkgPath)
	if err != nil {
		log.Fatal(err)
	}

	for key, imp := range pkg.Imports {
		keys := strings.Split(key, "/")
		packageName := keys[len(keys)-1]
		if packageName != expr.X.(*ast.Ident).Name {
			continue
		}

		s := imp.Types.Scope().Lookup(expr.Sel.Name)
		switch s.Type().Underlying().(type) {
		case *types.Interface:
			return true
		}
	}

	return false
}

func (r *ReceiverField) IsIdentMockable(ident *ast.Ident) bool {
	if ident.Obj != nil {
		switch ident.Obj.Decl.(*ast.TypeSpec).Type.(type) {
		case *ast.InterfaceType:
			return true
		}
	}

	return false
}
