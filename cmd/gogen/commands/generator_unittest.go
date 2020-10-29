package commands

import (
	"io"
	"log"

	"github.com/caudaganesh/go-generator/constant"
	"github.com/caudaganesh/go-generator/runner/unittestrunner"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createUnitTestCmd)
	addUnitTestFlags(createUnitTestCmd)
}

func addUnitTestFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("file", "f", "", "file path unit test generator")
	cmd.Flags().StringP("comment", "c", "", "comment for unit test generator")
	cmd.Flags().StringP("output", "o", "", "destination output of the result")
	cmd.MarkFlagRequired("file")
}

var createUnitTestCmd = &cobra.Command{
	Use:     "unittest",
	Short:   "Generate a unittest",
	Long:    `Generate a unittest`,
	Example: `gogen unittest -f="example/usecase/product.go" -o="example/usecase/product_test"`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		file, _ := cmd.Flags().GetString("file")
		comment, _ := cmd.Flags().GetString("comment")

		if file == "" {
			cmd.Usage()
			return
		}

		var out io.Reader
		var err error
		out, err = unittestrunner.Run(unittestrunner.UnitTestGenConf{
			File:    file,
			Comment: comment,
		})
		if err != nil {
			log.Fatal(err)
		}
		Write(output, out, constant.GenTest)
	},
}
