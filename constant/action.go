package constant

const (
	GenInterface = "GenerateInterface"
	GenProto     = "GenerateProto"
	GenUC        = "GenerateUC"
)

var MapActionsToPrefix = map[string]string{
	GenInterface: ".go",
	GenProto:     ".proto",
	GenUC:        ".go",
	"":           ".go",
}
