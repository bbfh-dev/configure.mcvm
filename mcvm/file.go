package mcvm

import (
	"path/filepath"

	"github.com/emersion/go-appdir"
)

// Returns the OS-specific path to an MCVM config file.
//
// e.g. ~/.config/mcvm/file.json
func ConfigFile(path ...string) string {
	dirs := appdir.New("mcvm")
	return filepath.Join(dirs.UserConfig(), filepath.Join(path...))
}

// Returns the OS-specific path to an MCVM data file.
//
// e.g. ~/.local/share/mcvm/file.json
func DataFile(path ...string) string {
	dirs := appdir.New("mcvm")
	return filepath.Join(dirs.UserData(), filepath.Join(path...))
}
