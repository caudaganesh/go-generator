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
	cmd.Flags().StringP("file", "f", "", "file path of the target struct")
	cmd.Flags().StringP("pkg", "p", "", "package of the target struct")
	cmd.Flags().StringP("comment", "c", "", "comment for the generated interface")
	cmd.Flags().StringP("name", "n", "", "name for the generated interface")
	cmd.Flags().StringP("target", "r", "", "target struct for interface generator")
	cmd.Flags().StringP("pkgName", "e", "", "package name for the generated interface")
	cmd.MarkFlagRequired("target")
	cmd.MarkFlagRequired("pkgName")
	cmd.MarkFlagRequired("name")
}

var createInterfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "Generate an interface",
	Long:  `Generate an interface`,
	Example: `
	gogen interface -e=usecase -f=example/usecase/product.go -r=ProductUC -n=ProductUseCase -c="ProductUseCase comments" -o=example/usecase/product_intf
	gogen interface -e=usecase -p=github.com/caudaganesh/go-generator/example/usecase -r=ProductUC -n=ProductUseCase -c="ProductUseCase comments" -o=example/usecase/product_intf
	`,
	Run: func(cmd *cobra.Command, args []string) {
		pkgName, _ := cmd.Flags().GetString("pkgName")
		pkg, _ := cmd.Flags().GetString("pkg")
		output, _ := cmd.Flags().GetString("output")
		file, _ := cmd.Flags().GetString("file")
		target, _ := cmd.Flags().GetString("target")
		name, _ := cmd.Flags().GetString("name")
		comment, _ := cmd.Flags().GetString("comment")

		var out io.Reader
		var err error
		out, err = interfacerunner.Run(interfacerunner.Conf{
			PackageName:  pkgName,
			File:         file,
			TargetStruct: target,
			Name:         name,
			Comment:      comment,
			Package:      pkg,
		})
		if err != nil {
			log.Fatal(err)
		}
		Write(output, out, constant.GenInterface)
	},
}
