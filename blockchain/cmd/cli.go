package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//type BoundedContext struct {
//	Blockchain *blockchain.BlockChain
//}

var rootCmd = &cobra.Command{
	Use:   "bc",
	Short: "Blockchain sandbox",
	Long: `Play around with type of blockchain implementation and
	how this is working under diffrent context and behaviour`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Execute root cmd has error", err)
		os.Exit(1)
	}
}
