package usecaserunner

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/caudaganesh/go-generator/config"
	"github.com/caudaganesh/go-generator/generator/applayer/usecase"
	"github.com/caudaganesh/go-generator/pkgloader"
	"github.com/caudaganesh/go-generator/structtype"
)

type UCGenConf struct {
	Package string
	Entity  string
}

func Run(conf UCGenConf) (io.Reader, error) {
	pkg, decls := pkgloader.LoadPackageDecls(conf.Package)
	str := structtype.GetFromDeclsByName(decls, conf.Entity)

	ucConf := config.GetUseCaseConfig()
	ucGen := usecase.NewUseCaseGen(
		usecase.Options{
			Package: conf.Package,
			Entity:  conf.Entity,
		},
		ucConf,
		pkg,
		str,
	)

	baseTemplate, err := ioutil.ReadFile(ucConf.TemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	bt := string(baseTemplate)
	result, err := ucGen.Generate(bt)

	return bytes.NewReader(result), err
}
