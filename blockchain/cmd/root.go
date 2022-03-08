package cmd

import (
	"fmt"
	"os"

	"github.com/raibru/blockchain/ledger"
	"github.com/raibru/blockchain/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var projectBase string
var userLicense string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config yaml file (default manifest.yaml)")
	//rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	//rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	//rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from current directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("./manifest.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	var cfg config.Config = config.Config{}
	cfg.DbPath = viper.GetString("database.path")
	cfg.DbFile = viper.GetString("database.file")
	cfg.Difficulty = viper.GetInt("blockchain.pow.difficulty")
	cfg.Root = viper.GetString("blockchain.block.root")
	config.SetConfig(&cfg)
	ledger.Init()
}
