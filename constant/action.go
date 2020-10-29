package constant

const (
	GenInterface = "GenerateInterface"
	GenProto     = "GenerateProto"
	GenUC        = "GenerateUC"
	GenTest      = "GenTest"
)

var MapActionsToPrefix = map[string]string{
	GenInterface: ".go",
	GenProto:     ".proto",
	GenUC:        ".go",
	GenTest:      ".go",
	"":           ".go",
}
