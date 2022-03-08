package cmd

import (
	"fmt"
	"os"

	"github.com/raibru/blockchain/ledger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the blockchain contents stored in database",
	Long:  "Print tho whole contents of the blockchain hold in database",
	Run: func(cmd *cobra.Command, args []string) {
		if !ledger.DBexists() {
			fmt.Println("Nothing to print. No initialized database found")
			os.Exit(1)
		}

		chain := ledger.CreateBlockChain()
		defer chain.Database.Close()

		iter := chain.Iterator()

		for {
			block := iter.Next()
			pck := ledger.CreateProof(block).Validate()
			fmt.Println(block.ToString(pck))
			if len(block.HashPrev) == 0 {
				break
			}
		}
	},
}
