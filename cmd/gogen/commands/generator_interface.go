package commands

import (
	"io"
	"log"

	"github.com/caudaganesh/go-generator/constant"
	"github.com/caudaganesh/go-generator/runner/interfacerunner"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createInterfaceCmd)
	addGlobalFlags(createInterfaceCmd)
	addInterfaceFlags(createInterfaceCmd)
}

func addInterfaceFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("file", "f", "", "file path unit test generator")
	cmd.Flags().StringP("comment", "c", "", "comment for unit test generator")
	cmd.Flags().StringP("name", "n", "", "name for interface generator")
	cmd.Flags().StringP("target", "r", "", "target struct for interface generator")
	cmd.Flags().StringP("pkgName", "p", "", "package name for the generated interface")
	cmd.MarkFlagRequired("target")
	cmd.MarkFlagRequired("pkgName")
	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("name")
}

var createInterfaceCmd = &cobra.Command{
	Use:     "interface",
	Short:   "Generate an interface",
	Long:    `Generate an interface`,
	Example: `gogen interface -p=usecase -f=example/usecase/product.go -r=ProductUC -n=ProductUseCase -c="ProductUseCase comments" -o=example/usecase/product_intf`,
	Run: func(cmd *cobra.Command, args []string) {
		pkgName, _ := cmd.Flags().GetString("pkgName")
		output, _ := cmd.Flags().GetString("output")
		file, _ := cmd.Flags().GetString("file")
		target, _ := cmd.Flags().GetString("target")
		name, _ := cmd.Flags().GetString("name")
		comment, _ := cmd.Flags().GetString("comment")

		var out io.Reader
		var err error
		out, err = interfacerunner.Run(interfacerunner.InterfaceGenConf{
			PackageName:  pkgName,
			File:         file,
			TargetStruct: target,
			Name:         name,
			Comment:      comment,
		})
		if err != nil {
			log.Fatal(err)
		}
		Write(output, out, constant.GenInterface)
	},
}
