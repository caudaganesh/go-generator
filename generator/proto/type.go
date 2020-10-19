package proto

import (
	"github.com/caudaganesh/go-generator/types"
)

const (
	String = "string"
	Double = "double"
	Int8   = "int8"
	Int16  = "int16"
	Int32  = "int32"
	Int64  = "int64"
	Bool   = "bool"
)

var mapPrimToPType = map[string]string{
	types.String:  String,
	types.Float32: Double,
	types.Float64: Double,
	types.Int:     Int64,
	types.Int8:    Int8,
	types.Int16:   Int16,
	types.Int32:   Int32,
	types.Int64:   Int64,
	types.Bool:    Bool,
	"time":        "google.protobuf.Timestamp",
}

func TransformTypeToPtype(name string) string {
	pType := mapPrimToPType[name]
	if pType != "" {
		return pType
	}

	return name
}
