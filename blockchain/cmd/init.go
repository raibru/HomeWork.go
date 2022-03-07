package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/raibru/HomeWork/blockchain/blockchain"
	"github.com/raibru/HomeWork/blockchain/config"

	"github.com/spf13/cobra"
)

var dbCre bool
var dbDel bool

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVar(&dbCre, "create-db", false, `create a new blockchain database inside directory
defined in configuration with one root element`)
	initCmd.Flags().BoolVar(&dbDel, "delete-db", false, `delete blockchain database with all data`)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init the database of the blockchain",
	Long:  "Initialize the blockchain database to directory or rmove old data collection",
	Run: func(cmd *cobra.Command, args []string) {
		if dbDel {
			if 0 < len(args) {
				cmd.Help()
				fmt.Println("No arguments allowed for delete-db")
				os.Exit(1)
			}
			if err := deleteDatabase(); err != nil {
				fmt.Println("Failed to delete database", err)
				os.Exit(1)
			}
		}
		if dbCre {
			if 0 < len(args) {
				cmd.Help()
				fmt.Println("No arguments allowed for create-db")
				os.Exit(1)
			}
			if err := createDatabase(); err != nil {
				fmt.Println("Failed to create database with root block", err)
				os.Exit(1)
			}
		}
		if !dbDel && !dbCre {
			cmd.Help()
			os.Exit(1)
		}
	},
}

func deleteDatabase() error {
	if blockchain.DBexists() {
		d, err := os.Open((*config.GetConfig()).DbPath)
		if err != nil {
			return err
		}
		defer d.Close()
		names, err := d.Readdirnames(-1)
		if err != nil {
			return err
		}
		for _, name := range names {
			err = os.RemoveAll(filepath.Join((*config.GetConfig()).DbPath, name))
			if err != nil {
				return err
			}
		}
		fmt.Printf("Database deleted in: %s\n", (*config.GetConfig()).DbPath)
	} else {
		fmt.Printf("Database not found in: %s\n", (*config.GetConfig()).DbPath)
	}
	return nil
}

func createDatabase() error {
	fmt.Println((*config.GetConfig()).ToString())

	if blockchain.DBexists() {
		return errors.New("Failed to create database. Database already exists")
	}
	chain := blockchain.CreateBlockChain()
	defer chain.Database.Close()
	return nil
}
