package interfacegen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	file := "../../example/usecase/product.go"
	opts := Options{
		File:         file,
		TargetStruct: "ProductUC",
		PackageName:  "usecase",
		Name:         "ProductUseCase",
		Comment:      "ProductUseCase holds product use case methods",
	}

	got, _ := Generate(opts)
	fmt.Println(string(got))
	want := `package usecase

import (
	"context"

	"github.com/caudaganesh/go-generator/example/entity"
)

// ProductUseCase holds product use case methods
type ProductUseCase interface {

	// GetProducts get list of product
	GetProducts(c context.Context) ([]entity.Product, error)
	// GetProduct get one product
	GetProduct(c context.Context) (entity.Product, error)
	// AddProduct add new product
	AddProduct(c context.Context, req entity.Product) error
	// UpdateProduct update existing product
	UpdateProduct(c context.Context, req entity.Product) error
	// DeleteProduct will delete existing product
	DeleteProduct(c context.Context, req entity.Product) error
}
`
	assert.Equal(t, want, string(got))
}
