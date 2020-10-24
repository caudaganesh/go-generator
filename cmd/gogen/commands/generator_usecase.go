package commands

import (
	"log"

	"github.com/caudaganesh/go-generator/constant"
	"github.com/caudaganesh/go-generator/runner/usecaserunner"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createUCCmd)
	addGlobalFlags(createUCCmd)
	addUCFlags(createUCCmd)
}

func addUCFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("pkg", "p", "", "package containing struct")
	cmd.Flags().StringP("target", "r", "", "target struct")
	cmd.MarkFlagRequired("pkg")
	cmd.MarkFlagRequired("target")
}

var createUCCmd = &cobra.Command{
	Use:     "usecase",
	Short:   "Generate a usecase",
	Long:    `Generate a usecase`,
	Example: `gogen usecase -p=github.com/caudaganesh/go-generator/example/entity -r=Product -o=example/usecase/product`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg, _ := cmd.Flags().GetString("pkg")
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Flags().GetString("output")

		out, err := usecaserunner.Run(usecaserunner.Conf{
			Package: pkg,
			Entity:  target,
		})
		if err != nil {
			log.Fatal(err)
		}

		Write(output, out, constant.GenUC)
	},
}
