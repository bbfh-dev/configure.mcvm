package screens

import (
	"github.com/bbfh-dev/configure.mcvm/ui/components"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var ProfilesScreen tea.Model

func SetupProfilesScreen(callback func(bool, components.SimpleListItem)) *tea.Model {
	ProfilesScreen = components.SimpleListComponent("HI", []list.Item{
		components.SimpleListItem("TEST1"),
	}, callback)
	return &ProfilesScreen
}
