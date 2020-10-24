package protorunner

import (
	"bytes"
	"io"

	"github.com/caudaganesh/go-generator/generator/proto"
)

type Conf struct {
	File         string
	TargetStruct string
	PackageName  string
	Name         string
	GoPackage    string
}

func Run(conf Conf) (io.Reader, error) {
	opt := proto.Options{
		File:         conf.File,
		TargetStruct: conf.TargetStruct,
		PackageName:  conf.PackageName,
		Name:         conf.Name,
		GoPackage:    conf.GoPackage,
	}
	result, err := proto.Generate(opt)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(result), err
}
