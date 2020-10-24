package repositoryrunner

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/caudaganesh/go-generator/config"
	"github.com/caudaganesh/go-generator/generator/applayer/repository"
	"github.com/caudaganesh/go-generator/pkgloader"
	"github.com/caudaganesh/go-generator/structtype"
)

type Conf struct {
	Package string
	Entity  string
}

func Run(conf Conf) (io.Reader, error) {
	pkg, decls := pkgloader.LoadPackageDecls(conf.Package)
	str := structtype.GetFromDeclsByName(decls, conf.Entity)

	cfg := config.GetRepositoryConfig()
	gen := repository.NewGen(
		repository.Options{
			Package: conf.Package,
			Entity:  conf.Entity,
		},
		cfg,
		pkg,
		str,
	)

	baseTemplate, err := ioutil.ReadFile(cfg.TemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	bt := string(baseTemplate)
	result, err := gen.Generate(bt)

	return bytes.NewReader(result), err
}
