package proto

import (
	"fmt"
	"testing"

	"github.com/caudaganesh/go-generator/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	file := "../../example/entity/product.go"

	opt := Options{
		File:         file,
		TargetStruct: "Product",
		PackageName:  "proto",
		Name:         "Product",
		GoPackage:    "./proto",
	}
	got, _ := Generate(opt)
	want := testhelper.GetExpectFromFile("./expect.txt")
	assert.Equal(t, want, string(got))
	fmt.Println(string(want))

}
