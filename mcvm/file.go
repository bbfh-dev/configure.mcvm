package mcvm

import (
	"path/filepath"

	"github.com/emersion/go-appdir"
)

// The mcvm.json configuration file
var MCVMConfig Config

// The ../internal/auth/db.json file
var MCVMAuth Auth

func GetConfigFile(name string) string {
	dirs := appdir.New("mcvm")
	return filepath.Join(dirs.UserConfig(), name)
}

func GetDataFile(path ...string) string {
	dirs := appdir.New("mcvm")
	return filepath.Join(dirs.UserData(), filepath.Join(path...))
}
