package types

import "go/ast"

const (
	String     = "string"
	Complex64  = "complex64"
	Complex128 = "complex128"
	Float32    = "float32"
	Float64    = "float64"
	Uint8      = "uint8"
	Uint16     = "uint16"
	Uint32     = "uint32"
	Uint64     = "uint64"
	Int        = "int"
	Int8       = "int8"
	Int16      = "int16"
	Int32      = "int32"
	Int64      = "int64"
	UintPtr    = "uintptr"
	Error      = "error"
	Bool       = "bool"
)

var primitives = []string{
	String,
	Complex64,
	Complex128,
	Float32,
	Float64,
	Uint8,
	Uint16,
	Uint32,
	Uint64,
	Int,
	Int8,
	Int16,
	Int32,
	Int64,
	UintPtr,
	Error,
	Bool,
}

func IsPrimitives(val string) bool {
	for _, primitive := range primitives {
		if val == primitive {
			return true
		}
	}

	return false
}

// GetPrimitiveType will return the primitive type if possible, if not just return the type
func GetPrimitiveType(typ string, obj *ast.Object) string {
	if obj == nil {
		return typ
	}

	if IsPrimitives(typ) {
		return typ
	}

	ts, ok := obj.Decl.(*ast.TypeSpec)
	if !ok {
		return typ
	}

	ident, ok := ts.Type.(*ast.Ident)
	if ok {
		return ident.Name
	}

	return typ
}
