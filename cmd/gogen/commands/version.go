package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gogen",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gogen version 1.0.0")
	},
}
