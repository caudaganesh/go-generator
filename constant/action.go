package constant

const (
	GenInterface = "GenerateInterface"
	GenProto     = "GenerateProto"
)

var MapActionsToPrefix = map[string]string{
	GenInterface: ".go",
	GenProto:     ".proto",
	"":           ".go",
}
