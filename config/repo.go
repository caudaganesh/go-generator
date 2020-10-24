package config

type RepositoryConfig struct {
	PackageName        string
	StructSuffix       string
	ContextPackageName string
	PlaceholderStyle   string
	ReferenceTag       string
	TemplatePath       string
	TableFormat        string
}

const (
	PlaceHolderStyleDollar       = "$"
	PlaceHolderStyleQuestionMark = "?"
)

func GetRepositoryConfig() RepositoryConfig {
	return RepositoryConfig{
		PackageName:        "repository",
		PlaceholderStyle:   PlaceHolderStyleDollar,
		ReferenceTag:       "db",
		ContextPackageName: "context",
		TableFormat:        "tbl_%s",
		StructSuffix:       "Repo",
		TemplatePath:       "example/template/repo.tmpl",
	}
}
