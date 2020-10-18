package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func addGlobalFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("output", "o", "", "destination output of the result")
}

var rootCmd = &cobra.Command{
	Use:   "gogen",
	Short: "gogen contains multiple go generator",
	Long:  `gogen contains multiple go generator : interface, proto`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
