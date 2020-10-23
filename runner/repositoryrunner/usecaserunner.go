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

type RepoGenConf struct {
	Package string
	Entity  string
}

func Run(conf RepoGenConf) (io.Reader, error) {
	pkg, decls := pkgloader.LoadPackageDecls(conf.Package)
	str := structtype.GetFromDeclsByName(decls, conf.Entity)

	repoConf := config.GetRepositoryConfig()
	ucGen := repository.NewRepoGen(
		repository.Options{
			Package: conf.Package,
			Entity:  conf.Entity,
		},
		repoConf,
		pkg,
		str,
	)

	baseTemplate, err := ioutil.ReadFile(repoConf.TemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	bt := string(baseTemplate)
	result, err := ucGen.Generate(bt)

	return bytes.NewReader(result), err
}
