package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// var section
var (
	major     = "0"
	minor     = "8"
	patch     = "1"
	buildTag  = "-"
	buildDate = "-"
	appName   = "bchain - blockchain structure playground"
	author    = "raibru <github.com/raibru>"
	license   = "MIT License (c) 2022 raibru"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bc",
	Long:  `The version of blockchain playground`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s - v%s.%s.%s\n", appName, major, minor, patch)
		fmt.Printf("  Build-%s (%s)\n", buildTag, buildDate)
		fmt.Printf("  author : %s\n", author)
		fmt.Printf("  license: %s\n\n", license)
	},
}
