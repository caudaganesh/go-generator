package config

type RepositoryConfig struct {
	PackageName        string
	StructSuffix       string
	ContextPackageName string
	TableDelimitation  uint8
	PlaceholderStyle   string
	AutoIncrField      string
	AutoIncrColumn     string
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
		TableDelimitation:  '_',
		PlaceholderStyle:   PlaceHolderStyleDollar,
		AutoIncrColumn:     "id",
		AutoIncrField:      "ID",
		ReferenceTag:       "db",
		ContextPackageName: "context",
		TableFormat:        "tbl_%s",
		StructSuffix:       "Repo",
		TemplatePath:       "example/template/repo.tmpl",
	}
}
