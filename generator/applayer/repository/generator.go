package repository

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"log"

	"github.com/caudaganesh/go-generator/config"
	"github.com/caudaganesh/go-generator/structtype"
	"github.com/iancoleman/strcase"
	"golang.org/x/tools/go/packages"
)

type Options struct {
	Package string
	Entity  string
}

type RepoGen struct {
	EntityPackage     string
	EntityPackageName string
	Entity            string
	EntityWithSpace   string
	TableName         string
	Struct            *structtype.StructType
	PropToTag         structtype.PropToTag
	PropToType        structtype.PropToType
	config.RepositoryConfig
}

func NewRepoGen(
	opt Options,
	repoConf config.RepositoryConfig,
	pkg *packages.Package,
	str *structtype.StructType,
) *RepoGen {

	return &RepoGen{
		EntityPackage:     pkg.ID,
		EntityPackageName: pkg.Types.Name(),
		Entity:            opt.Entity,
		EntityWithSpace:   strcase.ToDelimited(opt.Entity, ' '),
		TableName:         fmt.Sprintf(repoConf.TableFormat, strcase.ToSnake(opt.Entity)),
		Struct:            str,
		RepositoryConfig:  repoConf,
		PropToTag:         str.GetPropToTag(repoConf.ReferenceTag),
		PropToType:        str.GetPropToType(),
	}
}

// Generate takes in all of the fields and generate the repo
func (t *RepoGen) Generate(baseTemplate string) ([]byte, error) {
	tmpl := template.Must(template.New("repo").Parse(baseTemplate))
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, t)
	if err != nil {
		log.Fatal(err)
	}

	res := buf.String()
	res = html.UnescapeString(res)
	return []byte(res), nil
}
