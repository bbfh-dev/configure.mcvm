package screens

import (
	"github.com/bbfh-dev/configure.mcvm/ui/components"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var UsersScreen tea.Model

func SetupUsersScreen(callback func(bool, components.SimpleListItem)) *tea.Model {
	UsersScreen = components.SimpleListComponent("HI", []list.Item{
		components.SimpleListItem("TEST1"),
	}, callback)
	return &UsersScreen
}
