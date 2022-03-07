package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/raibru/HomeWork/blockchain/blockchain"
)

func init() {
	rootCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the blockchain contents stored in database",
	Long:  "Print tho whole contents of the blockchain hold in database",
	Run: func(cmd *cobra.Command, args []string) {
		if !blockchain.DBexists() {
			fmt.Println("Nothing to print. No initialized database found")
			os.Exit(1)
		}

		chain := blockchain.CreateBlockChain()
		defer chain.Database.Close()

		iter := chain.Iterator()

		for {
			block := iter.Next()
			pck := blockchain.CreateProof(block).Validate()
			fmt.Println(block.ToString(pck))
			if len(block.HashPrev) == 0 {
				break
			}
		}
	},
}
