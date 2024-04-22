package cli

import (
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
)

// CLI Flags passed to the program when called
var Flags struct {
	Log   string `short:"l" long:"log" description:"Path to log file"`
	Icons bool   `short:"i" long:"icons" description:"Enable icons in TUI"`
}

// Parses the CLI flags and sets default values
func ParseFlags() {
	_, err := flags.Parse(&Flags)
	if err != nil {
		os.Exit(0)
	}

	setDefaultFlags()
}

// Set the default flags
func setDefaultFlags() {
	if len(Flags.Log) <= 0 {
		Flags.Log = filepath.Join(os.TempDir(), "mcvmconf.log")
	}
}
