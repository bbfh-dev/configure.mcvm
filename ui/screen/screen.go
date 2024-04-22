package screen

import (
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Screen interface {
	Title() string
	Update(raw tea.Msg) (Screen, tea.Cmd)
	View(width int) []string
	Clear() Screen
	Messages() []tools.Message
	Notifications() []tools.Notification
	Keys() tools.Keybinds
	Lock() bool
}

func ContentHeight(screen Screen, width int) int {
	return lipgloss.Height(strings.Join(screen.View(width), "\n")) - 1
}

type ScreenTag string

const (
	HOME_SCREEN ScreenTag = "home"
	USER_SCREEN ScreenTag = "user"
)

type GoToScreen struct {
	Tag ScreenTag
}
