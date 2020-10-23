package usecase

import (
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

	pkg, decls := pkgloader.LoadPackageDecls(pkgPath)
	str := structtype.GetFromDeclsByName(decls, entity)
	ucConf := config.GetUseCaseConfig()
	ucConf.TemplatePath = "../../../example/template/usecase.tmpl"
	baseTemplate, err := ioutil.ReadFile(ucConf.TemplatePath)
	if err != nil {
		log.Fatal(err)
	}
	bt := string(baseTemplate)

	ucGen := NewUseCaseGen(
		Options{
			Package: pkgPath,
			Entity:  entity,
		},
		ucConf,
		pkg,
		str,
	)

	got, err := ucGen.Generate(bt)
	if err != nil {
		log.Fatal(err)
	}

	want := testhelper.GetExpectFromFile("./expect.txt")
	assert.Equal(t, want, string(got))
}
