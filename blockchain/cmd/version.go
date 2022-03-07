package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bc",
	Long:  `The version of blockchain playground`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bc blockchain playground v0.8.0")
	},
}
