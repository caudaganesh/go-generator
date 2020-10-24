package config

type DeliveryConfig struct {
	PackageName        string
	StructSuffix       string
	ContextPackageName string
	PlaceholderStyle   string
	ReferenceTag       string
	TemplatePath       string
	TableFormat        string
}

func GetDeliveryConfig() DeliveryConfig {
	return DeliveryConfig{
		PackageName:        "repository",
		ReferenceTag:       "json",
		ContextPackageName: "context",
		StructSuffix:       "Dlv",
		TemplatePath:       "example/template/dlv.tmpl",
	}
}
