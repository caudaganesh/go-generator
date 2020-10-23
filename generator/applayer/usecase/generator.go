package usecase

import (
	"bytes"
	"go/ast"
	"html"
	"html/template"
	"log"

	"github.com/caudaganesh/go-generator/config"
	"github.com/iancoleman/strcase"
	"golang.org/x/tools/go/packages"
)

type Options struct {
	Package string
	Entity  string
}

type UseCaseGen struct {
	EntityPackage     string
	EntityPackageName string
	Entity            string
	EntityWithSpace   string
	Struct            *ast.StructType
	config.UseCaseConfig
}

func NewUseCaseGen(
	opt Options,
	ucConf config.UseCaseConfig,
	pkg *packages.Package,
	str *ast.StructType,
) *UseCaseGen {

	return &UseCaseGen{
		EntityPackage:     pkg.ID,
		EntityPackageName: pkg.Types.Name(),
		Entity:            opt.Entity,
		EntityWithSpace:   strcase.ToDelimited(opt.Entity, ' '),
		Struct:            str,
		UseCaseConfig:     ucConf,
	}
}

// Generate takes in all of the fields and generate the use case
func (t *UseCaseGen) Generate(baseTemplate string) ([]byte, error) {
	tmpl := template.Must(template.New("usecase").Parse(baseTemplate))
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, t)
	if err != nil {
		log.Fatal(err)
	}

	res := buf.String()
	res = html.UnescapeString(res)
	return []byte(res), nil
}
