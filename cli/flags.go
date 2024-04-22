package cli

import (
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
)

var Flags struct {
	Log   string `short:"l" long:"log" description:"Path to log file"`
	Icons bool   `short:"i" long:"icons" description:"Enable icons in TUI"`
}

func ParseFlags() {
	_, err := flags.Parse(&Flags)
	if err != nil {
		os.Exit(0)
	}

	setDefaultFlags()
}

func setDefaultFlags() {
	if len(Flags.Log) <= 0 {
		Flags.Log = filepath.Join(os.TempDir(), "mcvmconf.log")
	}
}
