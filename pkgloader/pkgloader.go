package pkgloader

import (
	"fmt"
	"go/ast"
	"log"
	"sync"

	"golang.org/x/tools/go/packages"
)

var pkg sync.Map

// Load loads a packages by pkgPath
func Load(pkgPath string) (*packages.Package, error) {
	v, found := pkg.Load(pkgPath)
	if found {
		return v.(*packages.Package), nil
	}

	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedTypes | packages.NeedImports}
	p, err := packages.Load(cfg, pkgPath)
	if err != nil {
		return nil, err
	}
	if len(p) < 1 {
		return nil, fmt.Errorf("package not found")
	}
	pkg.Store(pkgPath, p[0])
	return p[0], nil
}

func LoadPackageDecls(pkgPath string) (*packages.Package, []ast.Decl) {
	pkg, err := Load(pkgPath)
	if err != nil {
		log.Fatal(err)
	}

	var decls []ast.Decl
	for _, f := range pkg.Syntax {
		decls = append(decls, f.Decls...)
	}

	return pkg, decls
}
