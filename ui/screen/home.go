package screen

import (
	"github.com/bbfh-dev/configure.mcvm/ui/component"
	"github.com/bbfh-dev/configure.mcvm/ui/style"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/bbfh-dev/configure.mcvm/ui/widget"
	tea "github.com/charmbracelet/bubbletea"
)

const WELCOME_MESSAGE = "Welcome! You can see available actions on the bottom of the screen, press `?` to toggle the help message."

type HomeScreen struct {
	title         string
	width         int
	height        int
	notifications []tools.Notification
	messages      []tools.Message

	radioWidget widget.RadioWidget
}

func NewHomeScreen() HomeScreen {
	return HomeScreen{
		title: tools.WithIcon(tools.HOME_ICON, "Manage MCVM Config"),
		radioWidget: widget.NewRadioWidget(component.NewLiteralList([]component.Item{
			component.NewLiteralItem(
				"users",
				tools.WithIcon(tools.USER_ICON, "Users"),
				"User management & default user",
			),
			component.NewLiteralItem(
				"profiles",
				tools.WithIcon(tools.PROFILE_ICON, "Profiles"),
				"Manage profiles & instances",
			),
			component.NewLiteralItem(
				"packages",
				tools.WithIcon(tools.PACKAGE_ICON, "Packages"),
				"Manage global packages",
			),
			component.NewLiteralItem(
				"preferences",
				tools.WithIcon(tools.GEAR_ICON, "Preferences"),
				"Manage preferences for how the whole program will work",
			),
			component.NewLiteralItem(
				"game_options",
				tools.WithIcon(tools.GAME_ICON, "Game options"),
				"Manage global options for both client and server that are inherited across all of your profiles",
			),
		}), true),
	}
}

func (screen HomeScreen) Title() string {
	return screen.title
}

func (screen HomeScreen) Update(raw tea.Msg) (Screen, tea.Cmd) {
	var commands []tea.Cmd

	switch msg := raw.(type) {

	case tea.WindowSizeMsg:
		screen.width = msg.Width
		screen.height = msg.Height
		screen.radioWidget.Width = msg.Width
	}

	screen.radioWidget = tools.UpdateModel[widget.RadioWidget](
		&commands,
		screen.radioWidget.Clear(),
		raw,
	)
	screen.notifications = append(screen.notifications, screen.radioWidget.Notification)

	for _, message := range screen.radioWidget.Messages {
		switch msg := message.(type) {
		case component.SelectItem:
			switch msg.Item {
			case "users":
				screen.messages = append(screen.messages, GoToScreen{Tag: USER_SCREEN})
			}
		}
	}

	return screen, tea.Batch(commands...)
}

func (screen HomeScreen) View(width int) []string {
	var contents []string

	tools.AppendContent(&contents, style.Info.RenderLine(width, WELCOME_MESSAGE), width)
	tools.AppendContent(&contents, screen.radioWidget.View(), width)

	return contents
}

func (screen HomeScreen) Clear() Screen {
	screen.messages = make([]tools.Message, 0)
	screen.notifications = make([]tools.Notification, 0)

	return screen
}

func (screen HomeScreen) Notifications() []tools.Notification {
	return screen.notifications
}

func (screen HomeScreen) Messages() []tools.Message {
	return screen.messages
}

func (screen HomeScreen) Keys() tools.Keybinds {
	return screen.radioWidget.Keys
}

func (screen HomeScreen) Lock() bool {
	return false
}
