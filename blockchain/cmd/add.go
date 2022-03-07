package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/raibru/HomeWork/blockchain/blockchain"
)

var blockContent string

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&blockContent, "block", "m", "", "add a block with data")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Block into blockchain with contents parameter",
	Long:  "Add a block to existing blockchain contents stored in database",
	Run: func(cmd *cobra.Command, args []string) {
		if !blockchain.DBexists() {
			fmt.Println("Nothing to add. No initialized database found")
			os.Exit(1)
		}

		if len(blockContent) > 1 {
			chain := blockchain.CreateBlockChain()
			defer chain.Database.Close()
			chain.AddBlock([]byte(blockContent))
			fmt.Println("Added Block!")
		}
	},
}
