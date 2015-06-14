package config

import (
	"fmt"
	"os"
)

// Config represents a configuration.
type Config struct {
	debug bool // print debug/verbose output
}

// New creates a Config.
func New(debug bool) *Config {
	return &Config{
		debug: debug,
	}
}

// Debug formats and outputs s to Stdout if the debug flag is turned on.
func (c *Config) Debug(s string, args ...interface{}) {
	if c == nil || !c.debug {
		return
	}
	fmt.Fprintf(os.Stdout, s, args...)
}
