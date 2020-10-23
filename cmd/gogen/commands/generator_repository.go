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
	cmd.Flags().StringP("pkg", "p", "", "package to entity")
	cmd.Flags().StringP("entity", "e", "", "target entity")
	cmd.MarkFlagRequired("pkg")
	cmd.MarkFlagRequired("entity")
}

var createRepoCmd = &cobra.Command{
	Use:     "repo",
	Short:   "Generate a repo",
	Long:    `Generate a repo`,
	Example: `gogen repo -p=github.com/caudaganesh/go-generator/example/entity -e=Product -o=example/repo/product`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg, _ := cmd.Flags().GetString("pkg")
		entity, _ := cmd.Flags().GetString("entity")
		output, _ := cmd.Flags().GetString("output")

		out, err := repositoryrunner.Run(repositoryrunner.RepoGenConf{
			Package: pkg,
			Entity:  entity,
		})
		if err != nil {
			log.Fatal(err)
		}

		Write(output, out, constant.GenUC)
	},
}
