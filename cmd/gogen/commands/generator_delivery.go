package commands

import (
	"log"

	"github.com/caudaganesh/go-generator/constant"
	"github.com/caudaganesh/go-generator/runner/deliveryrunner"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createDlvCmd)
	addGlobalFlags(createDlvCmd)
	addDlvCmd(createDlvCmd)
}

func addDlvCmd(cmd *cobra.Command) {
	cmd.Flags().StringP("pkg", "p", "", "package containing struct")
	cmd.Flags().StringP("target", "r", "", "target struct")
	cmd.MarkFlagRequired("pkg")
	cmd.MarkFlagRequired("target")
}

var createDlvCmd = &cobra.Command{
	Use:     "delivery",
	Short:   "Generate a delivery",
	Long:    `Generate a delivery`,
	Example: `gogen delivery -p=github.com/caudaganesh/go-generator/example/entity -r=Product -o=example/delivery/product`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg, _ := cmd.Flags().GetString("pkg")
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Flags().GetString("output")

		out, err := deliveryrunner.Run(deliveryrunner.Conf{
			Package: pkg,
			Entity:  target,
		})
		if err != nil {
			log.Fatal(err)
		}

		Write(output, out, constant.GenUC)
	},
}
