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
	repoConf := config.GetRepositoryConfig()

	pkg, decls := pkgloader.LoadPackageDecls(pkgPath)
	str := structtype.GetFromDeclsByName(decls, entity)

	repoConf.TemplatePath = "../../../example/template/repo.tmpl"
	baseTemplate, err := ioutil.ReadFile(repoConf.TemplatePath)
	if err != nil {
		log.Fatal(err)
	}
	bt := string(baseTemplate)

	ucGen := NewRepoGen(
		Options{
			Package: pkgPath,
			Entity:  entity,
		},
		repoConf,
		pkg,
		str,
	)

	got, err := ucGen.Generate(bt)
	if err != nil {
		log.Fatal(err)
	}

	want := testhelper.GetExpectFromFile("./expect.txt")
	fmt.Println(string(got))
	assert.Equal(t, want, string(got))
}
