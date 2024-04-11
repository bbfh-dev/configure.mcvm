package screens

import (
	"github.com/bbfh-dev/configure.mcvm/ui/components"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var IndexScreen tea.Model

func SetupIndexScreen(callback func(bool, components.SimpleListItem)) tea.Model {
	IndexScreen = components.SimpleListComponent("Choose what to configure:", []list.Item{
		components.SimpleListItem("Users"),
		components.SimpleListItem("Profiles"),
		components.SimpleListItem("Packages"),
		components.SimpleListItem("Preferences"),
		components.SimpleListItem("Global game options"),
	}, callback)
	return IndexScreen
}
