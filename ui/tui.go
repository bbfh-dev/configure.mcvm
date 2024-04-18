package ui

import (
	"github.com/bbfh-dev/configure.mcvm/ui/screens"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/bbfh-dev/configure.mcvm/ui/widgets"
	"github.com/charmbracelet/bubbles/key"
)

func NewIndexScreen() screens.IndexScreen {
	return screens.IndexScreen{
		TitleModel: widgets.NewTitleWidget(
			tools.WithIcon(tools.GEAR_ICON, "MCVM Configuration"),
		),
		ChooseConfigModel: widgets.NewRadioWidget(
			tools.WithIcon(tools.LIST_ICON, "Choose what to configure:"),
			tools.NewLiteralList(
				tools.NewLiteralItem("users",
					tools.WithIcon(tools.USER_ICON, "Users"),
					"User management & default user"),
				tools.NewLiteralItem("profiles",
					tools.WithIcon(tools.PROFILE_ICON, "Profiles"),
					"Manage profiles & instances"),
				tools.NewLiteralItem("packages",
					tools.WithIcon(tools.PACKAGE_ICON, "Packages"),
					"Manage global packages"),
				tools.NewLiteralItem("preferences",
					tools.WithIcon(tools.GEAR_ICON, "Preferences"),
					"Manage preferences for how the whole program will work"),
				tools.NewLiteralItem(
					"game_options",
					tools.WithIcon(tools.GAME_ICON, "Game options"),
					"Manage global options for both client and server that are inherited across all of your profiles",
				),
			),
		),
		HelpModel: widgets.NewHelpWidget(
			// Prefixed with <number>+ for sorting in the menu
			tools.ExtendWithDefaultKeybinds(tools.Keybinds{
				"0+scroll_up": key.NewBinding(
					key.WithKeys("K", "up"),
					key.WithHelp("K/󱕑", "scroll up"),
				),
				"0+scroll_down": key.NewBinding(
					key.WithKeys("J", "down"),
					key.WithHelp("J/󱕐", "scroll down"),
				),
			}),
		),
		CurrentScreen: screens.INDEX_SCREEN,
		Screens:       map[string]screens.Screen{},
	}
}
