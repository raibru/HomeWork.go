package config

import (
	"bytes"
	"fmt"
)

// Config hold the configuration data used in blockchain
type Config struct {
	DbPath     string
	DbFile     string
	Difficulty int
	Root       string
}

var cfg Config = Config{}

// SetConfig set the application configurations
func SetConfig(c *Config) {
	cfg = *c
}

// GetConfig return the application configurations
func GetConfig() *Config {
	return &cfg
}

// ToString of type Config answer a formated output of current config data
func (c Config) ToString() string {
	var out bytes.Buffer
	fmt.Fprintln(&out, "\nBlockchain configuration:")
	fmt.Fprintf(&out, "\tDatabase Path  : %s\n", c.DbPath)
	fmt.Fprintf(&out, "\tDatabase File  : %s\n", c.DbFile)
	fmt.Fprintf(&out, "\tDifficulty     : %d\n", c.Difficulty)
	fmt.Fprintf(&out, "\tRoot Block Data: %s\n\n", c.Root)
	return out.String()
}
