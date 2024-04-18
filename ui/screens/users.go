package screens

import (
	"fmt"
	"strings"

	"github.com/bbfh-dev/configure.mcvm/mcvm"
	"github.com/bbfh-dev/configure.mcvm/mcvm/schema"
	user_screen "github.com/bbfh-dev/configure.mcvm/ui/screens/user"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/bbfh-dev/configure.mcvm/ui/widgets"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UsersScreen struct {
	initialized    bool
	exitingToIndex bool
	Width          int
	height         int
	scroll         int
	maxScroll      int
	TitleModel     widgets.TitleWidget
	ListModel      widgets.RadioWidget
	Keys           tools.Keybinds
	notifications  []string
}

func NewUserListScreen() UsersScreen {
	return UsersScreen{
		TitleModel: widgets.NewTitleWidget(tools.WithIcon(tools.USER_ICON, "Manage users")),
		Keys: tools.Keybinds{
			"z-go-back": key.NewBinding(
				key.WithKeys("esc"),
				key.WithHelp(tools.IconFallback("󱊷 ", "esc"), "go back"),
			),
		},
		ListModel: widgets.NewRadioWidget(
			tools.WithIcon(tools.LIST_ICON, "User list:"),
			user_screen.NewUserList(),
		),
	}
}

func (screen UsersScreen) Init() tea.Cmd {
	return nil
}

func (screen UsersScreen) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var commands []tea.Cmd
	var saveNotification string

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, screen.Keys["z-go-back"]):
			screen.exitingToIndex = true
			return screen, nil
		}
	}

	var model tea.Model
	var cmd tea.Cmd

	// List Model
	model, cmd = screen.ListModel.Update(message)
	screen.ListModel = model.(widgets.RadioWidget)
	commands = append(commands, cmd)

	if screen.ListModel.HasSelected() {
		switch screen.ListModel.Selection {
		case "new":
			mcvm.MCVMConfig.Users["unnamed-user"] = schema.User{
				Type: new(string),
			}
		case "save":
			// TODO: Handle errors
			mcvm.MCVMConfig.EncodeToFile()
			saveNotification = fmt.Sprintf("Saved to disk: %s", mcvm.GetConfigFile("mcvm.json"))
		}
		screen.ListModel.Selection = ""
	}

	// Notifications
	screen.notifications = []string{
		screen.ListModel.Notification,
		saveNotification,
	}

	return screen, nil
}

func (screen UsersScreen) View() string {
	return ""
}

func (screen UsersScreen) Render() (string, []string, *int) {
	header := screen.TitleModel.View()
	contents := screen.RenderContents()
	scroll := &screen.scroll
	return header, contents, scroll
}

func (screen UsersScreen) SetWidth(width int) Screen {
	screen.Width = width
	screen.TitleModel.Width = width
	return screen
}

func (screen UsersScreen) Height() int {
	return screen.height - lipgloss.Height(screen.TitleModel.View())
}

func (screen UsersScreen) SetHeight(height int) Screen {
	screen.height = height
	return screen
}

func (screen UsersScreen) Keybinds() tools.Keybinds {
	return screen.Keys
}

func (screen UsersScreen) DidExit() bool {
	return screen.exitingToIndex
}

func (screen UsersScreen) Exit() bool {
	if screen.exitingToIndex {
		screen.ListModel.Blur()
		return true
	}

	return false
}

func (screen UsersScreen) Reload() Screen {
	screen.exitingToIndex = false
	screen.ListModel = screen.ListModel.Focus()

	for key, value := range screen.ListModel.Keys {
		screen.Keys[key] = value
	}

	return screen
}

func (screen UsersScreen) Notifications() []string {
	return screen.notifications
}

func (screen UsersScreen) Scroll() int {
	return screen.scroll
}

func (screen UsersScreen) SetScroll(scroll int) Screen {
	screen.scroll = scroll
	return screen
}

func (screen UsersScreen) MaxScroll() int {
	return screen.maxScroll
}

func (screen UsersScreen) SetMaxScroll(scroll int) Screen {
	screen.maxScroll = scroll
	return screen
}

func (screen UsersScreen) RenderContents() []string {
	contents := []string{""}
	contents = append(contents, strings.Split(screen.ListModel.View(), "\n")...)

	return append(contents, "")
}
