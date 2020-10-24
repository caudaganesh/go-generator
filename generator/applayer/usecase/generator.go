package usecase

import (
	"bytes"
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

type Gen struct {
	EntityPackage     string
	EntityPackageName string
	Entity            string
	EntityWithSpace   string
	Struct            *structtype.StructType
	config.UseCaseConfig
}

func NewGen(
	opt Options,
	ucConf config.UseCaseConfig,
	pkg *packages.Package,
	str *structtype.StructType,
) *Gen {

	return &Gen{
		EntityPackage:     pkg.ID,
		EntityPackageName: pkg.Types.Name(),
		Entity:            opt.Entity,
		EntityWithSpace:   strcase.ToDelimited(opt.Entity, ' '),
		Struct:            str,
		UseCaseConfig:     ucConf,
	}
}

// Generate takes in all of the fields and generate the use case
func (g *Gen) Generate(baseTemplate string) ([]byte, error) {
	tmpl := template.Must(template.New("usecase").Parse(baseTemplate))
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, g)
	if err != nil {
		log.Fatal(err)
	}

	res := buf.String()
	res = html.UnescapeString(res)
	return []byte(res), nil
}
