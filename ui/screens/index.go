package screens

import (
	"slices"
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/styles"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/bbfh-dev/configure.mcvm/ui/widgets"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const MINIMUM_WIDTH = 32
const MINIMUM_HEIGHT = 10

const WELCOME_MESSAGE = "Welcome! You can see available actions on the bottom of the screen, press `?` to toggle the help message."

type CurrentScreen string

const (
	INDEX_SCREEN     CurrentScreen = "index"
	USER_LIST_SCREEN CurrentScreen = "users"
)

type IndexScreen struct {
	initialized       bool
	Width             int
	Height            int
	Scroll            int
	MaxScroll         int
	TitleModel        widgets.TitleWidget
	ChooseConfigModel widgets.RadioWidget
	HelpModel         widgets.HelpWidget
	CurrentScreen     CurrentScreen
	Screens           map[string]Screen
	notifications     []string
}

func (screen IndexScreen) Init() tea.Cmd {
	return nil
}

func (screen IndexScreen) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if !screen.initialized {
		screen = screen.Initialize()
	}

	var commands []tea.Cmd

	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		screen = screen.UpdateScreenSizes(msg.Width, msg.Height)
		screen.TitleModel.Width = msg.Width
		screen.ChooseConfigModel.Width = msg.Width
		screen.ChooseConfigModel.Height = msg.Height
		screen.Width = msg.Width
		screen.Height = msg.Height
		if screen.Height != 0 {
			if screen.CurrentScreen == INDEX_SCREEN {
				screen.Scroll = max(0, screen.Scroll-msg.Height-screen.Height)
			} else {
				screen.Screens[string(screen.CurrentScreen)] = screen.Screens[string(screen.CurrentScreen)].SetScroll(
					max(0, screen.Screens[string(screen.CurrentScreen)].Scroll()-msg.Height-screen.Height),
				)
			}
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, screen.HelpModel.Keys["quit"]):
			return screen, tea.Quit
		case key.Matches(msg, screen.HelpModel.Keys["0+scroll_up"]):
			if screen.CurrentScreen == INDEX_SCREEN {
				if screen.Scroll > 0 {
					screen.Scroll -= 1
				}
			} else if screen.Screens[string(screen.CurrentScreen)].Scroll() > 0 {
				screen.Screens[string(screen.CurrentScreen)] = screen.Screens[string(screen.CurrentScreen)].SetScroll(
					screen.Screens[string(screen.CurrentScreen)].Scroll() - 1,
				)
			}
			return screen, nil
		case key.Matches(msg, screen.HelpModel.Keys["0+scroll_down"]):
			if screen.CurrentScreen == INDEX_SCREEN {
				height := screen.Height
				height -= lipgloss.Height(screen.TitleModel.View())
				height -= lipgloss.Height(screen.HelpModel.View())
				screen.MaxScroll = len(screen.RenderContents()) - height + 1 + max(0, lipgloss.Height(screen.RenderNotifications())-2)
				if screen.Scroll < screen.MaxScroll {
					screen.Scroll += 1
				}
			} else {
				height := screen.Screens[string(screen.CurrentScreen)].Height()
				height -= lipgloss.Height(screen.HelpModel.View())
				screen.Screens[string(screen.CurrentScreen)] = screen.Screens[string(screen.CurrentScreen)].SetMaxScroll(
					len(screen.Screens[string(screen.CurrentScreen)].RenderContents()) - height + max(0, lipgloss.Height(screen.RenderNotifications())-2),
				)
				if screen.Screens[string(screen.CurrentScreen)].Scroll() < screen.Screens[string(screen.CurrentScreen)].MaxScroll() {
					screen.Screens[string(screen.CurrentScreen)] = screen.Screens[string(screen.CurrentScreen)].SetScroll(
						screen.Screens[string(screen.CurrentScreen)].Scroll() + 1,
					)
				}
			}
			return screen, nil
		}
	}

	var model tea.Model
	var cmd tea.Cmd

	// Help Model
	model, cmd = screen.HelpModel.Update(message)
	screen.HelpModel = model.(widgets.HelpWidget)
	commands = append(commands, cmd)

	// Update screens
	switch screen.CurrentScreen {
	case USER_LIST_SCREEN:
		model, cmd = screen.Screens[string(USER_LIST_SCREEN)].Update(message)
		screen.Screens[string(USER_LIST_SCREEN)] = model.(Screen)
		commands = append(commands, cmd)

		if screen.Screens[string(USER_LIST_SCREEN)].Exit() {
			screen.CurrentScreen = INDEX_SCREEN
		}
	default:
		// Choose Config Model
		model, cmd = screen.ChooseConfigModel.Update(message)
		screen.ChooseConfigModel = model.(widgets.RadioWidget)
		commands = append(commands, cmd)

		if screen.ChooseConfigModel.HasSelected() {
			switch screen.ChooseConfigModel.Selection {
			case string(USER_LIST_SCREEN):
				screen.CurrentScreen = USER_LIST_SCREEN
				if _, ok := screen.Screens[string(USER_LIST_SCREEN)]; !ok {
					screen.Screens[string(USER_LIST_SCREEN)] = NewUserListScreen()
				}
				screen.ChooseConfigModel.Notification = ""
				screen.Screens[string(USER_LIST_SCREEN)] = screen.Screens[string(USER_LIST_SCREEN)].Reload()
				screen = screen.UpdateScreenSizes(screen.Width, screen.Height)
				screen = screen.FillHelpFor(string(USER_LIST_SCREEN))
			}
		}
	}

	// Notifications
	screen.notifications = []string{
		screen.ChooseConfigModel.Notification,
	}
	for _, model := range screen.Screens {
		if model.DidExit() {
			continue
		}

		screen.notifications = append(screen.notifications, model.Notifications()...)
	}

	return screen, nil
}

func (screen IndexScreen) View() string {
	if screen.Width < MINIMUM_WIDTH || screen.Height < MINIMUM_HEIGHT {
		return "Window is too small!" + strings.Repeat("\n", max(0, screen.Height-2))
	}

	var scroll *int
	var header string
	var contents []string
	footer := screen.RenderNotifications()
	footer += screen.HelpModel.View()

	switch screen.CurrentScreen {
	case USER_LIST_SCREEN:
		model, ok := screen.Screens["users"]
		if !ok {
			// TODO: Handle users screen not existing
		}
		header, contents, scroll = model.Render()
	default:
		header, contents, scroll = screen.PopulateHelp().Render()
	}

	return RenderScreen(
		CalculateScrollbar(
			screen.Height-lipgloss.Height(header)-lipgloss.Height(footer),
			len(contents),
			scroll,
		),
		&screen.Width,
		&screen.Height,
		header,
		footer,
		contents,
	)
}

func (screen IndexScreen) Render() (string, []string, *int) {
	header := screen.TitleModel.View()
	contents := screen.RenderContents()
	scroll := &screen.Scroll
	return header, contents, scroll
}

func (screen IndexScreen) FillHelpFor(modelName string) IndexScreen {
	for key := range screen.HelpModel.Keys {
		if !slices.Contains(screen.HelpModel.DefaultKeys, key) {
			delete(screen.HelpModel.Keys, key)
		}
	}

	screen.HelpModel.Id = modelName
	for key, value := range screen.Screens[modelName].Keybinds() {
		screen.HelpModel.Keys[key] = value
	}

	return screen
}

func (screen IndexScreen) PopulateHelp() IndexScreen {
	if screen.HelpModel.Id == "index" {
		return screen
	}

	screen.HelpModel.Id = "index"
	for key, value := range screen.ChooseConfigModel.Keys {
		screen.HelpModel.Keys[key] = value
	}

	return screen
}

func (screen IndexScreen) RenderContents() []string {
	contents := []string{""}
	contents = append(
		contents,
		strings.Split(styles.Info.Render(screen.Width-2, WELCOME_MESSAGE), "\n")...,
	)
	contents = append(contents, "")
	contents = append(contents, strings.Split(screen.ChooseConfigModel.View(), "\n")...)

	return append(contents, "")
}

func (screen IndexScreen) RenderNotifications() string {
	var contents string

	for _, notification := range screen.notifications {
		if len(notification) == 0 {
			continue
		}

		contents += styles.Notification.Render(
			screen.Width,
			tools.WithIcon(tools.INFO_ICON, notification),
		) + "\n"
	}

	return contents
}

func (screen IndexScreen) Initialize() IndexScreen {
	screen = screen.PopulateHelp()
	screen.ChooseConfigModel = screen.ChooseConfigModel.Focus()
	screen.initialized = true

	return screen
}

func (screen IndexScreen) UpdateScreenSizes(width int, height int) IndexScreen {
	for key, model := range screen.Screens {
		screen.Screens[key] = model.SetWidth(width).SetHeight(height)
	}

	return screen
}
