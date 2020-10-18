package interfacerunner

import (
	"bytes"
	"io"

	"github.com/caudaganesh/go-generator/generator/interfacegen"
)

type InterfaceGenConf struct {
	PackageName  string
	File         string
	TargetStruct string
	Name         string
	Comment      string
}

func Run(conf InterfaceGenConf) (io.Reader, error) {
	opt := interfacegen.Options{
		File:         conf.File,
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
