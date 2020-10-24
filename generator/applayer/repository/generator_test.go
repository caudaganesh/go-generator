package repository

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/caudaganesh/go-generator/config"
	"github.com/caudaganesh/go-generator/pkgloader"
	"github.com/caudaganesh/go-generator/structtype"
	"github.com/caudaganesh/go-generator/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	pkgPath := "github.com/caudaganesh/go-generator/example/entity"
	entity := "Product"
	cfg := config.GetRepositoryConfig()

	pkg, decls := pkgloader.LoadPackageDecls(pkgPath)
	str := structtype.GetFromDeclsByName(decls, entity)

	cfg.TemplatePath = "../../../example/template/repo.tmpl"
	baseTemplate, err := ioutil.ReadFile(cfg.TemplatePath)
	if err != nil {
		log.Fatal(err)
	}
	bt := string(baseTemplate)

	gen := NewGen(
		Options{
			Package: pkgPath,
			Entity:  entity,
		},
		cfg,
		pkg,
		str,
	)

	got, err := gen.Generate(bt)
	if err != nil {
		log.Fatal(err)
	}

	want := testhelper.GetExpectFromFile("./expect.txt")
	fmt.Println(string(got))
	assert.Equal(t, want, string(got))
}
