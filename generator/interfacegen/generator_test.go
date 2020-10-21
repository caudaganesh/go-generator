package interfacegen

import (
	"testing"

	"github.com/caudaganesh/go-generator/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestGenerateWithPackage(t *testing.T) {
	opts := Options{
		Package:      "github.com/caudaganesh/go-generator/example/usecase",
		TargetStruct: "ProductUC",
		PackageName:  "usecase",
		Name:         "ProductUseCase",
		Comment:      "ProductUseCase holds product use case methods",
	}

	got, _ := Generate(opts)

	want := testhelper.GetExpectFromFile("./expect.txt")
	assert.Equal(t, want, string(got))
}

func TestGenerateWithFile(t *testing.T) {
	opts := Options{
		File:         "../../example/usecase/product.go",
		TargetStruct: "ProductUC",
		PackageName:  "usecase",
		Name:         "ProductUseCase",
		Comment:      "ProductUseCase holds product use case methods",
	}

	got, _ := Generate(opts)
	want := testhelper.GetExpectFromFile("./expect.txt")
	assert.Equal(t, want, string(got))
}
