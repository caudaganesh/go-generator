package pkgloader

import (
	"sync"

	"golang.org/x/tools/go/packages"
)

var pkgList sync.Map

// Load loads a package by pkgPath
func Load(pkgPath string) (*packages.Package, error) {
	v, found := pkgList.Load(pkgPath)
	if found {
		return v.(*packages.Package), nil
	}

	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes}
	pkg, err := packages.Load(cfg, pkgPath)
	return pkg[0], err
}
