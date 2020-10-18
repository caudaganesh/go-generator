package tool

var primitives = []string{
	"string",
	"complex64",
	"complex128",
	"float32",
	"float64",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uintptr",
	"error",
	"bool",
}

func IsPrimitives(val string) bool {
	for _, primitive := range primitives {
		if val == primitive {
			return true
		}
	}

	return false
}
