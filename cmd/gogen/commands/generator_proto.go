package commands

import (
	"log"

	"github.com/caudaganesh/go-generator/constant"
	"github.com/caudaganesh/go-generator/runner/protorunner"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createProtoCmd)
	addProtoFlags(createProtoCmd)
}

func addProtoFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("file", "f", "", "file path to target struct")
	cmd.Flags().StringP("target", "r", "", "target struct")
	cmd.Flags().StringP("pkgName", "p", "", "package name for generated proto")
	cmd.Flags().StringP("name", "n", "", "message name for generated proto")
	cmd.Flags().StringP("goPkg", "g", "", "go package for generated proto")
	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("name")
}

var createProtoCmd = &cobra.Command{
	Use:     "proto",
	Short:   "Generate a proto",
	Long:    `Generate a proto`,
	Example: `gogen proto -p="product" -r="Product" -n="Product" -f="example/entity/product.go" -o="example/proto/product" -g="proto"`,
	Run: func(cmd *cobra.Command, args []string) {
		pkgName, _ := cmd.Flags().GetString("pkgName")
		file, _ := cmd.Flags().GetString("file")
		target, _ := cmd.Flags().GetString("target")
		name, _ := cmd.Flags().GetString("name")
		goPkg, _ := cmd.Flags().GetString("goPkg")
		output, _ := cmd.Flags().GetString("output")

		out, err := protorunner.Run(protorunner.ProtoGenConf{
			File:         file,
			TargetStruct: target,
			PackageName:  pkgName,
			Name:         name,
			GoPackage:    goPkg,
		})
		if err != nil {
			log.Fatal(err)
		}
		Write(output, out, constant.GenProto)
	},
}
