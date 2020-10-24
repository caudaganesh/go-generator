package delivery

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
	cfg := config.GetDeliveryConfig()

	pkg, decls := pkgloader.LoadPackageDecls(pkgPath)
	str := structtype.GetFromDeclsByName(decls, entity)

	cfg.TemplatePath = "../../../example/template/dlv.tmpl"
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

	fmt.Println(string(got))
	want := testhelper.GetExpectFromFile("./expect.txt")
	assert.Equal(t, want, string(got))
}
