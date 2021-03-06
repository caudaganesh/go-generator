package commands

import (
	"log"

	"github.com/caudaganesh/go-generator/constant"
	"github.com/caudaganesh/go-generator/runner/repositoryrunner"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createRepoCmd)
	addGlobalFlags(createRepoCmd)
	addRepoCmd(createRepoCmd)
}

func addRepoCmd(cmd *cobra.Command) {
	cmd.Flags().StringP("pkg", "p", "", "package containing struct")
	cmd.Flags().StringP("target", "r", "", "target struct")
	cmd.MarkFlagRequired("pkg")
	cmd.MarkFlagRequired("target")
}

var createRepoCmd = &cobra.Command{
	Use:     "repo",
	Short:   "Generate a repo",
	Long:    `Generate a repo`,
	Example: `gogen repo -p=github.com/caudaganesh/go-generator/example/entity -r=Product -o=example/repo/product`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg, _ := cmd.Flags().GetString("pkg")
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Flags().GetString("output")

		out, err := repositoryrunner.Run(repositoryrunner.Conf{
			Package: pkg,
			Entity:  target,
		})
		if err != nil {
			log.Fatal(err)
		}

		Write(output, out, constant.GenUC)
	},
}
