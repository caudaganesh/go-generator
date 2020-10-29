package unittest

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"github.com/caudaganesh/go-generator/pkgloader"
)

// CalledExpr is a extension of  `ast.CallExpr`
type CalledExpr struct {
	*ast.CallExpr
	fn      *Func
	decls   []ast.Decl
	pkgPath string
}

// Name get the name of called
func (c *CalledExpr) Name() string {
	if se, ok := c.Fun.(*ast.SelectorExpr); ok {
		return se.Sel.Name
	}
	return ""
}

// GetArgsReturns return the args and result of the `CalledExpr`
func (c *CalledExpr) GetArgsReturns() (args, returns []string, err error) {
	if c.fn.Signature.Recv == nil {
		return nil, nil, fmt.Errorf("only support a method")
	}

	se, ok := c.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, nil, fmt.Errorf("failed to assert selector expr")
	}

	var fieldName string
	// After getting the method name, we need to get the field name
	// then get the package path
	if se, ok := se.X.(*ast.SelectorExpr); ok {
		fieldName = se.Sel.Name
	}

	// return immediately if we're unable to get the fieldName
	if fieldName == "" {
		return nil, nil, fmt.Errorf("unsupported method without field")
	}

	for _, field := range c.fn.Struct.Fields.List {
		if len(field.Names) < 1 || field.Names[0].Name != fieldName {
			continue
		}

		identifier, interfaceName := c.getIntfNameAndIdent(field)
		if identifier == nil {
			continue
		}

		pkgPath := c.getPkgPath(identifier.Name)
		if pkgPath == "" {
			return nil, nil, fmt.Errorf("package path for %s is not found", identifier.Name)
		}

		p, err := pkgloader.Load(pkgPath)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot load package %v, because %v", identifier.Name, err)
		}

		fun := p.Types.Scope().Lookup(interfaceName)
		if fun == nil {
			return nil, nil, fmt.Errorf("unresolved import, %s", se.Sel.Name)
		}

		itf, ok := fun.Type().Underlying().(*types.Interface)
		if !ok {
			return nil, nil, fmt.Errorf("%s is not an interface", interfaceName)
		}

		args, returns = c.ParseIntfToArgsAndReturns(itf, se.Sel.Name)
		return args, returns, nil
	}

	return nil, nil, nil
}

func (c *CalledExpr) getIntfNameAndIdent(field *ast.Field) (*ast.Ident, string) {
	var interfaceName string
	var identifier *ast.Ident
	switch i := field.Type.(type) {
	case *ast.Ident:
		identifier = i
		interfaceName = i.Name
	case *ast.SelectorExpr:
		identifier = i.X.(*ast.Ident)
		interfaceName = i.Sel.Name
	}
	return identifier, interfaceName
}

func (c *CalledExpr) getPkgPath(pkgName string) string {
	pkgPath := c.pkgPath
	for _, decl := range c.decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.IMPORT {
			continue
		}

		for _, spec := range gd.Specs {
			ip, ok := spec.(*ast.ImportSpec)
			if !ok {
				continue
			}

			pPath := strings.Trim(ip.Path.Value, "\"")
			pPaths := strings.Split(pPath, "/")
			if pPaths[len(pPaths)-1] == pkgName {
				pkgPath = pPath
				break
			}
		}
	}

	return pkgPath
}

func (c *CalledExpr) ParseIntfToArgsAndReturns(itf *types.Interface, fnName string) (args, returns []string) {
	for i := 0; i < itf.NumMethods(); i++ {
		method := itf.Method(i)
		if method.Name() != fnName {
			continue
		}

		signature, ok := method.Type().Underlying().(*types.Signature)
		if !ok {
			continue
		}

		args, returns = findDataType(signature.Params()), findDataType(signature.Results())
	}

	return
}

// FieldName returns the filed name of the `CalledExpr`
// in the struct receiver
func (c *CalledExpr) FieldName() string {
	if se, ok := c.Fun.(*ast.SelectorExpr); ok {
		if se, ok := se.X.(*ast.SelectorExpr); ok {
			return se.Sel.Name
		}
	}
	return ""
}

// findDataType finds the list of datatype from a `types.Tuple`
func findDataType(t *types.Tuple) (res []string) {
	for i := 0; i < t.Len(); i++ {
		paths := strings.Split(t.At(i).Type().String(), "/")
		res = append(res, paths[len(paths)-1])
	}
	return
}
