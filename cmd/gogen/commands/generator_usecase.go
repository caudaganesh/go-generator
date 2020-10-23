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
	cmd.Flags().StringP("pkg", "p", "", "package to entity")
	cmd.Flags().StringP("entity", "e", "", "target entity")
	cmd.MarkFlagRequired("pkg")
	cmd.MarkFlagRequired("entity")
}

var createUCCmd = &cobra.Command{
	Use:     "usecase",
	Short:   "Generate a usecase",
	Long:    `Generate a usecase`,
	Example: `gogen usecase -p=github.com/caudaganesh/go-generator/example/entity -e=Product -o=example/usecase/product`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg, _ := cmd.Flags().GetString("pkg")
		entity, _ := cmd.Flags().GetString("entity")
		output, _ := cmd.Flags().GetString("output")

		out, err := usecaserunner.Run(usecaserunner.UCGenConf{
			Package: pkg,
			Entity:  entity,
		})
		if err != nil {
			log.Fatal(err)
		}

		Write(output, out, constant.GenUC)
	},
}
