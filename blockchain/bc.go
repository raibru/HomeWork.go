package main

// see Link: https://dev.to/nheindev/building-a-blockchain-in-go-pt-iii-persistence-3884

import (
	"os"

	"github.com/raibru/blockchain/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
