package tools

import (
	"fmt"

	"github.com/bbfh-dev/configure.mcvm/cli"
)

type icon string

// https://www.nerdfonts.com/cheat-sheet
const (
	GEAR_ICON     icon = " "
	INFO_ICON     icon = " "
	USER_ICON     icon = " "
	PROFILE_ICON  icon = "󱃶 "
	PACKAGE_ICON  icon = " "
	GAME_ICON     icon = "󰍳 "
	LIST_ICON     icon = "󱉯 "
	SAVE_ICON     icon = "󰆓 "
	HOME_ICON     icon = " "
	NEW_USER_ICON icon = " "
	WARNING_ICON  icon = " "
	LOADING_ICON  icon = " "
	DELETE_ICON   icon = " "
	STAR_ICON     icon = " "
)

func WithIcon(icon icon, str string) string {
	if cli.Flags.Icons {
		return fmt.Sprintf("%s %s", icon, str)
	}

	return str
}

func IconFallback(icon icon, fallback string) string {
	if cli.Flags.Icons {
		return string(icon)
	}

	return fallback
}
