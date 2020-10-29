package unittestrunner

import (
	"bytes"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/caudaganesh/go-generator/generator/unittest"
	"github.com/caudaganesh/go-generator/generator/unittest/template"
	"github.com/caudaganesh/go-generator/pkgloader"
)

type UnitTestGenConf struct {
	File    string
	Comment string
}

func Run(conf UnitTestGenConf) (io.Reader, error) {
	fs := token.NewFileSet()
	file, _ := parser.ParseFile(fs, conf.File, nil, 0)
	packageName := file.Name.String()
	dir := filepath.Dir(conf.File)
	wd, _ := os.Getwd()
	dir = wd + "/" + dir
	pkgPath, err := pkgloader.ParsePackageImport(dir)
	if err != nil {
		log.Fatal(err)
	}
	funcs := unittest.GetFunctionsFromFile(file, pkgPath)
	opt := unittest.Options{
		Comment:        conf.Comment,
		Template:       template.NewGomockTemplate(),
		HeaderTemplate: template.NewHeaderTemplate(),
		PackageName:    packageName,
		Funcs:          funcs,
	}
	generator := unittest.NewGenerator(opt)
	buf := bytes.NewBuffer([]byte{})
	err = generator.Write(buf)

	return bytes.NewReader(buf.Bytes()), err
}
