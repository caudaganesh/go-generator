package template

func NewHeaderTemplate() string {
	return `
	package {{.Package}}
	
	import (
		"testing"
	
		"github.com/stretchr/testify/assert"
	)
	{{if .Comment}}
	//{{.Comment}}
	{{end}}`
}
