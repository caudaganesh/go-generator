package delivery

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

type Gen struct {
	EntityPackage     string
	EntityPackageName string
	Entity            string
	EntityWithSpace   string
	TableName         string
	Struct            *structtype.StructType
	PropToTag         structtype.PropToTag
	PropToType        structtype.PropToType
	PropToCamel       structtype.PropToCamel
	config.DeliveryConfig
}

func NewGen(
	opt Options,
	cfg config.DeliveryConfig,
	pkg *packages.Package,
	str *structtype.StructType,
) *Gen {

	return &Gen{
		EntityPackage:     pkg.ID,
		EntityPackageName: pkg.Types.Name(),
		Entity:            opt.Entity,
		EntityWithSpace:   strcase.ToDelimited(opt.Entity, ' '),
		TableName:         fmt.Sprintf(cfg.TableFormat, strcase.ToSnake(opt.Entity)),
		Struct:            str,
		DeliveryConfig:    cfg,
		PropToTag:         str.GetPropToTag(cfg.ReferenceTag),
		PropToType:        str.GetPropToType(),
		PropToCamel:       str.GetPropToCamel(),
	}
}

// Generate takes in all of the fields and generate the repo
func (g *Gen) Generate(baseTemplate string) ([]byte, error) {
	tmpl := template.Must(template.New("delivery").Parse(baseTemplate))
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, g)
	if err != nil {
		log.Fatal(err)
	}

	res := buf.String()
	res = html.UnescapeString(res)
	return []byte(res), nil
}
