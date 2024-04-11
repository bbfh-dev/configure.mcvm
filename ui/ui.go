package ui

import (
	"github.com/bbfh-dev/configure.mcvm/ui/components"
	"github.com/bbfh-dev/configure.mcvm/ui/screens"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	IsIndexScreen    uint8 = 0
	IsUsersScreen    uint8 = 1
	IsProfilesScreen uint8 = 2
)

var currentScreen uint8

type model struct {
	width  int
	height int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	switch currentScreen {
	case IsIndexScreen:
		screens.IndexScreen, cmd = screens.IndexScreen.Update(msg)
	case IsUsersScreen:
		screens.UsersScreen, cmd = screens.UsersScreen.Update(msg)
	case IsProfilesScreen:
		screens.ProfilesScreen, cmd = screens.ProfilesScreen.Update(msg)
	default:
		cmd = nil
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	switch currentScreen {
	case IsUsersScreen:
		return screens.UsersScreen.View()
	case IsProfilesScreen:
		return screens.ProfilesScreen.View()
	default:
		return screens.IndexScreen.View()
	}
}

func MainModel() model {
	screens.SetupIndexScreen(func(exit bool, item components.SimpleListItem) {
		if exit {
			return
		}

		defaultCallback := func(exit bool, item components.SimpleListItem) {
			if exit {
				currentScreen = IsIndexScreen
			}
		}

		switch item {
		case "Users":
			screens.SetupUsersScreen(defaultCallback)
			currentScreen = IsUsersScreen
		case "Profiles":
			screens.SetupProfilesScreen(defaultCallback)
			currentScreen = IsProfilesScreen
		}
	})

	return model{}
}
