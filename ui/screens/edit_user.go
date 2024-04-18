package screens

import (
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/bbfh-dev/configure.mcvm/ui/widgets"
	tea "github.com/charmbracelet/bubbletea"
)

type UserEditScreen struct {
	initialized    bool
	exitingToIndex bool
	Width          int
	height         int
	scroll         int
	maxScroll      int
	TitleModel     widgets.TitleWidget
	Keys           tools.Keybinds
	notifications  []string
}

func (screen UserEditScreen) Init() tea.Cmd {
	return nil
}

func (screen UserEditScreen) Update(messages tea.Msg) (tea.Model, tea.Cmd) {
	return screen, nil
}

func (screen UserEditScreen) View() string {
	return ""
}

func (screen UserEditScreen) SetWidth(width int) Screen {
	screen.Width = width
	return screen
}

func (screen UserEditScreen) Height() int {
	return screen.height
}

func (screen UserEditScreen) SetHeight(height int) Screen {
	screen.height = height
	return screen
}

func (screen UserEditScreen) Scroll() int {
	return screen.scroll
}

func (screen UserEditScreen) MaxScroll() int {
	return screen.maxScroll
}

func (screen UserEditScreen) SetScroll(scroll int) Screen {
	screen.scroll = scroll
	return screen
}

func (screen UserEditScreen) SetMaxScroll(scroll int) Screen {
	screen.maxScroll = scroll
	return screen
}

func (screen UserEditScreen) Render() (string, []string, *int) {
	return "", []string{}, new(int)
}

func (screen UserEditScreen) Keybinds() tools.Keybinds {
	return tools.Keybinds{}
}

func (screen UserEditScreen) DidExit() bool {
	return screen.exitingToIndex
}

func (screen UserEditScreen) Exit() bool {
	return screen.exitingToIndex
}

func (screen UserEditScreen) Reload() Screen {
	return screen
}

func (screen UserEditScreen) Notifications() []string {
	return []string{}
}

func (screen UserEditScreen) RenderContents() []string {
	return []string{}
}
