package mcvm

import (
	"path/filepath"

	"github.com/emersion/go-appdir"
)

func ConfigFile(path ...string) string {
	dirs := appdir.New("mcvm")
	return filepath.Join(dirs.UserConfig(), filepath.Join(path...))
}

func DataFile(path ...string) string {
	dirs := appdir.New("mcvm")
	return filepath.Join(dirs.UserData(), filepath.Join(path...))
}
