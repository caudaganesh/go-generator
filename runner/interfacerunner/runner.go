package interfacerunner

import (
	"bytes"
	"io"

	"github.com/caudaganesh/go-generator/generator/interfacegen"
)

type Conf struct {
	PackageName  string
	File         string
	Package      string
	TargetStruct string
	Name         string
	Comment      string
}

func Run(conf Conf) (io.Reader, error) {
	opt := interfacegen.Options{
		File:         conf.File,
		Package:      conf.Package,
		TargetStruct: conf.TargetStruct,
		PackageName:  conf.PackageName,
		Name:         conf.Name,
		Comment:      conf.Comment,
	}
	result, err := interfacegen.Generate(opt)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(result), err
}
