package unittest

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"path/filepath"
	"testing"

	"github.com/caudaganesh/go-generator/generator/unittest/template"
	"github.com/caudaganesh/go-generator/pkgloader"
)

func TestNewGenerator(t *testing.T) {
	t.Run("generate", func(t *testing.T) {
		filePath := "../../example/usecase/product.go"
		fs := token.NewFileSet()
		dir := filepath.Dir(filePath)
		pkgPath, _ := pkgloader.ParsePackageImport(dir)
		file, _ := parser.ParseFile(fs, filePath, nil, 0)
		packageName := file.Name.String()
		funcs := GetFunctionsFromFile(file, pkgPath)

		generator := NewGenerator(Options{
			PackageName:    packageName,
			Funcs:          funcs,
			Template:       template.NewGomockTemplate(),
			HeaderTemplate: template.NewHeaderTemplate(),
			Comment:        "Testing Comment",
		})

		buf := bytes.NewBuffer([]byte{})
		generator.Write(buf)

		fmt.Println(buf.String())
	})
}
