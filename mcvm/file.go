package mcvm

import (
	"path/filepath"

	"github.com/emersion/go-appdir"
)

// The mcvm.json configuration file
var MCVMConfig Config

func GetMCVMFile(name string) string {
	dirs := appdir.New("mcvm")
	return filepath.Join(dirs.UserConfig(), name)
}
