package utils

import (
	"fmt"

	"github.com/bbfh-dev/configure.mcvm/cli"
)

type Icon string

// https://www.nerdfonts.com/cheat-sheet
const (
	GEAR_ICON = " "
	INFO_ICON = " "
)

func WithIcon(icon Icon, str string) string {
	if cli.Flags.Icons {
		return fmt.Sprintf("%s %s", icon, str)
	}

	return str
}
