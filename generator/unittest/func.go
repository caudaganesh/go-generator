package unittest

import (
	"fmt"
	"go/ast"
	"go/token"
)

//Func is a wrapper around ast.FuncDecl containing few methods
//to use within a test template
type Func struct {
	PkgDecls  []ast.Decl
	File      *ast.File
	Struct    *ast.StructType
	Signature *ast.FuncDecl
	PkgPath   string
}

//NewFunc returns pointer to the Func struct
func NewFunc(
	sig *ast.FuncDecl,
	str *ast.StructType,
	file *ast.File,
	pkgDecls []ast.Decl,
	pkgPath string) *Func {
	return &Func{
		Signature: sig,
		Struct:    str,
		File:      file,
		PkgDecls:  pkgDecls,
		PkgPath:   pkgPath,
	}
}

//NumParams returns a number of the function params
func (f *Func) NumParams() int {
	return f.Signature.Type.Params.NumFields()
}

//NumResults returns a number of the function results
func (f *Func) NumResults() int {
	if f.Signature.Type.Results == nil {
		return 0
	}
	return f.Signature.Type.Results.NumFields()
}

//Params returns a list of the function params with their types
func (f *Func) Params(fs *token.FileSet) []string {
	if f.Signature.Type.Params == nil {
		return nil
	}

	params := []string{}
	for i, p := range f.Signature.Type.Params.List {
		for _, n := range p.Names {
			param := f.ParseFieldsToNodeString(fs, n, f.Signature.Type.Params.List, p, i)
			params = append(params, param)
		}
	}

	return params
}

//Fields a list of the function fields with their types
func (f *Func) Fields(fs *token.FileSet) []string {
	fields := []string{}
	if f.Signature.Recv == nil {
		return fields
	}

	for i := range f.Signature.Recv.List {
		for _, fl := range f.Struct.Fields.List {
			for _, n := range fl.Names {
				param := f.ParseFieldsToNodeString(fs, n, f.Struct.Fields.List, fl, i)
				fields = append(fields, param)
			}
		}
	}

	return fields
}

func (f *Func) ParseFieldsToNodeString(
	fs *token.FileSet,
	ident *ast.Ident,
	fields []*ast.Field,
	field *ast.Field,
	index int) (nodeString string) {

	nodeString = ident.Name
	if index == len(fields)-1 && f.IsVariadic() {
		nodeString += " []" + nodeToString(fs, field.Type.(*ast.Ellipsis).Elt)
		return
	}

	nodeString += " " + nodeToString(fs, field.Type)
	return
}

func (f *Func) ReceiverFields() []ReceiverField {
	if f.Signature.Recv.List == nil {
		return nil
	}

	fields := []ReceiverField{}
	for _, v := range f.Signature.Recv.List {
		switch xv := v.Type.(type) {
		case *ast.StarExpr:
			fields = f.ParseStarExprToReceiverField(xv)
		case *ast.Ident:
		}
	}

	return fields
}

func (f *Func) ParseStarExprToReceiverField(se *ast.StarExpr) []ReceiverField {
	st := f.Struct
	if st == nil {
		return nil
	}

	var fields []ReceiverField
	calledMethods := f.FindCalledMethods()
	for _, fst := range st.Fields.List {
		for _, n := range fst.Names {
			name := n.Name
			field, ok := n.Obj.Decl.(*ast.Field)
			if !ok {
				continue
			}
			receiverField := NewReceiverField(name, f.PkgPath, field, calledMethods)
			receiverField.Set()
			fields = append(fields, receiverField)
		}
	}

	return fields
}

func (f *Func) HasReceiverFields() bool {
	return f.IsMethod() && len(f.ReceiverFields()) > 0
}

func (f *Func) HasMockableField(fields []ReceiverField) bool {
	for _, f := range fields {
		if f.Mockable {
			return true
		}
	}

	return false
}

//Results returns a list of the function results with their types
//if function's last param is an error it is not included in the result slice
func (f *Func) Results(fs *token.FileSet) []string {
	if f.Signature.Type.Results == nil {
		return nil
	}

	var (
		results []string
		n       = 1
	)
	for _, r := range f.Signature.Type.Results.List {
		if len(r.Names) <= 0 {
			results = append(results, fmt.Sprintf("got%d %s", n, nodeToString(fs, r.Type)))
			n++
			continue
		}

		for range r.Names {
			results = append(results, fmt.Sprintf("got%d %s", n, nodeToString(fs, r.Type)))
			n++
		}
	}

	if f.ReturnsError() {
		results = results[:len(results)-1]
	}

	return results
}

//ParamsNames returns a list of the function params' names
func (f *Func) ParamsNames() []string {
	if f.Signature.Type.Params == nil {
		return nil
	}

	names := []string{}
	for i, p := range f.Signature.Type.Params.List {
		for k, n := range p.Names {
			name := n.Name
			if i == len(f.Signature.Type.Params.List)-1 && k == len(p.Names)-1 && f.IsVariadic() {
				name += "..."
			}
			names = append(names, name)
		}
	}

	return names
}

//ResultsNames returns a list of the function results' names
//if function's last result is an error the name of param is "err"
func (f *Func) ResultsNames() []string {
	if f.Signature.Type.Results == nil {
		return nil
	}

	var (
		names []string
		n     = 1
	)

	for _, r := range f.Signature.Type.Results.List {
		if len(r.Names) <= 0 {
			names = append(names, fmt.Sprintf("got%d", n))
			n++
			continue
		}

		for range r.Names {
			names = append(names, fmt.Sprintf("got%d", n))
			n++
		}
	}

	if f.ReturnsError() {
		names[len(names)-1] = "err"
	}

	return names
}

//Name returns a name of func
func (f *Func) Name() string {
	return f.Signature.Name.String()
}

//TestName returns a name of the test
func (f *Func) TestName() string {
	name := "Test"
	if f.IsMethod() {
		recvType := f.ReceiverType()
		if star, ok := recvType.(*ast.StarExpr); ok {
			name += star.X.(*ast.Ident).String()
		} else {
			name += recvType.(*ast.Ident).String()
		}
		name += "_"
	} else if !f.Signature.Name.IsExported() {
		name += "_"
	}

	return name + f.Signature.Name.String()
}

//IsMethod returns true if the function is a method
func (f *Func) IsMethod() bool {
	return f.Signature.Recv != nil
}

//ReceiverType returns a type of the method receiver
func (f *Func) ReceiverType() ast.Expr {
	if f.Signature.Recv == nil {
		return nil
	}
	return f.Signature.Recv.List[0].Type
}

//ReceiverInstance returns a instance of the method receiver
func (f *Func) ReceiverInstance() string {
	if f.Signature.Recv == nil {
		return ""
	}

	switch ri := f.Signature.Recv.List[0].Type.(type) {
	case *ast.StarExpr:
		if si, ok := ri.X.(*ast.Ident); ok {
			return "&" + si.Name
		}
	case *ast.Ident:
		return ri.Name
	}

	return ""
}

//ReturnsError returns true if the function's last param's type is error
func (f *Func) ReturnsError() bool {
	lastResult := f.LastResult()
	if lastResult == nil {
		return false
	}

	ident, ok := lastResult.Type.(*ast.Ident)
	return ok && ident.Name == "error"
}

//LastParam returns function's last param
func (f *Func) LastParam() *ast.Field {
	numFields := len(f.Signature.Type.Params.List)
	if numFields == 0 {
		return nil
	}

	return f.Signature.Type.Params.List[numFields-1]
}

//LastResult returns function's last result
func (f *Func) LastResult() *ast.Field {
	if f.Signature.Type.Results == nil {
		return nil
	}

	numFields := len(f.Signature.Type.Results.List)
	if numFields == 0 {
		return nil
	}

	return f.Signature.Type.Results.List[numFields-1]
}

//IsVariadic returns true if it's the variadic function
func (f *Func) IsVariadic() bool {
	lastParam := f.LastParam()
	if lastParam == nil {
		return false
	}

	_, isVariadic := lastParam.Type.(*ast.Ellipsis)

	return isVariadic
}

// FindCalledMethods find all called method inside the function body
func (f *Func) FindCalledMethods() (fns []*CalledExpr) {
	for _, stmt := range f.Signature.Body.List {
		as, ok := stmt.(*ast.AssignStmt)
		if !ok || len(as.Rhs) < 1 { // ensure the Right Hand Statement is at least 1 and a method
			continue
		}

		ce, ok := as.Rhs[0].(*ast.CallExpr)
		if !ok {
			continue
		}

		fns = append(fns, &CalledExpr{
			CallExpr: ce,
			fn:       f,
			pkgPath:  f.PkgPath,
			decls:    f.PkgDecls,
		})
	}

	return
}
