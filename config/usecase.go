package config

type UseCaseConfig struct {
	PackageName        string
	StructSuffix       string
	ContextPackageName string
	TemplatePath       string
}

func GetUseCaseConfig() UseCaseConfig {
	return UseCaseConfig{
		PackageName:        "usecase",
		StructSuffix:       "UC",
		ContextPackageName: "context",
		TemplatePath:       "example/template/usecase.tmpl", //TODO: change this to absolute path to be able run anywhere
	}
}
