package proto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	file := "../../example/entity/product.go"

	opt := Options{
		File:         file,
		TargetStruct: "Product",
		PackageName:  "productproto",
		Name:         "Product",
		GoPackage:    "./productproto",
	}
	got, _ := Generate(opt)
	want := `
	syntax="proto3";

	package productproto;

	option go_package = ./productproto;

	message Product {
		
		int64 id = 1;
		string name = 2;
		bool active = 3;
		double discount = 4;
		int64 price = 5;
		google.protobuf.Timestamp created_at = 6;
		int64 product_status = 7;
	}
	`
	assert.Equal(t, want, string(got))
	fmt.Println(string(got))
}