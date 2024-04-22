package ui

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbfh-dev/configure.mcvm/mcvm"
	"github.com/bbfh-dev/configure.mcvm/mcvm/auth"
	"github.com/bbfh-dev/configure.mcvm/mcvm/config"
	"github.com/bbfh-dev/configure.mcvm/ui/screen"
	"github.com/bbfh-dev/configure.mcvm/ui/style"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/bbfh-dev/configure.mcvm/ui/widget"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// The height of `style.Title.Render()`
const TITLE_HEIGHT = 2

const MINIMUM_WIDTH = 32
const MINIMUM_HEIGHT = 10

// ---- IndexModel

type IndexModel struct {
	width         int
	height        int
	scroll        int
	keys          tools.Keybinds
	Screens       map[screen.ScreenTag]screen.Screen
	Current       screen.ScreenTag
	notifications map[string]tools.Notification

	helpWidget widget.HelpWidget
}

func NewIndexModel() IndexModel {
	return IndexModel{
		keys: tools.GlobalKeybinds,
		Screens: map[screen.ScreenTag]screen.Screen{
			screen.HOME_SCREEN: screen.NewHomeScreen(),
		},
		Current:    screen.HOME_SCREEN,
		helpWidget: widget.NewHelpWidget(tools.GlobalKeybinds),
		notifications: map[string]tools.Notification{
			"load+mcvm.json": {
				Type: tools.TASK,
				Text: "Loading mcvm.json...",
			},
		},
	}
}

func (model IndexModel) Init() tea.Cmd {
	return tea.Batch(loadConfig, loadAuth)
}

func (model IndexModel) Update(raw tea.Msg) (tea.Model, tea.Cmd) {
	var commands []tea.Cmd

	switch msg := raw.(type) {

	case tea.WindowSizeMsg:
		model.width = msg.Width
		model.height = msg.Height

	case tea.KeyMsg:
		if model.Screens[model.Current].Lock() {
			if msg.String() == "ctrl+c" {
				return model, tea.Quit
			}
			break
		}
		switch {

		case key.Matches(msg, tools.QuitKeybind):
			return model, tea.Quit

		case key.Matches(msg, tools.GlobalKeybinds["scroll.up"]):
			model.scroll -= 1

		case key.Matches(msg, tools.GlobalKeybinds["scroll.down"]):
			model.scroll += 1

		case key.Matches(msg, tools.GlobalKeybinds["goto.top"]):
			model.scroll = 0

		case key.Matches(msg, tools.GlobalKeybinds["goto.bottom"]):
			model.scroll = math.MaxInt
		}

	case ErrTaskMsg:
		delete(model.notifications, msg.Task)
		model.notifications["error"] = tools.Notification{
			Type: tools.ERROR,
			Text: fmt.Sprintf("ERROR(%s): %s", msg.Task, msg.Err.Error()),
		}

	case EndTaskMsg:
		delete(model.notifications, msg.Task)
	}

	_, space := model.getSpace()
	model = model.limitScroll(space)

	// Screen
	var cmd tea.Cmd
	model.Screens[model.Current], cmd = model.Screens[model.Current].Clear().Update(raw)
	commands = append(commands, cmd)

	for _, message := range model.Screens[model.Current].Messages() {
		switch msg := message.(type) {
		case screen.GoToScreen:
			model.Current = msg.Tag
			if _, ok := model.Screens[model.Current]; !ok {
				model.Screens[model.Current] = setupScreen(model.Current)
			}

			return model.Update(nil)

		case widget.SaveMsg:
			model.notifications["save+mcvm.json"] = tools.Notification{
				Type: tools.TASK,
				Text: "Writing changes to mcvm.json...",
			}
			model.Current = screen.HOME_SCREEN
			commands = append(commands, saveConfig)
		}
	}

	// Widgets
	model.helpWidget = tools.UpdateModel[widget.HelpWidget](&commands, model.helpWidget, raw)

	model.helpWidget.Keys = tools.GlobalKeybinds
	for key, bind := range model.Screens[model.Current].Keys() {
		model.helpWidget.Keys[key] = bind
	}

	return model, tea.Batch(commands...)
}

func (model IndexModel) View() string {
	if model.height < MINIMUM_HEIGHT || model.width < MINIMUM_WIDTH {
		return "Window is too small!" + strings.Repeat("\n", max(0, model.height-2))
	}

	var builder strings.Builder
	footer, space := model.getSpace()
	contentHeight := screen.ContentHeight(model.Screens[model.Current], model.width-1)

	builder.WriteString(style.Title.RenderLine(model.width, model.Screens[model.Current].Title()))

	contents := model.Screens[model.Current].View(model.width - 1)
	window := contents[max(0, model.scroll):max(0, min(len(contents)-1, space+model.scroll))]
	barHeight, barPosition := tools.ScrollBar(
		space,
		contentHeight,
		model.scroll,
	)

	if contentHeight <= space {
		for _, line := range window {
			builder.WriteString(line + "\n")
		}
	} else {
		for i, line := range window {
			if i >= barPosition && i < barHeight+barPosition {
				builder.WriteString(line + style.ScrollForeground + "\n")
			} else {
				builder.WriteString(line + style.ScrollBackground + "\n")
			}
		}
	}

	builder.WriteString(tools.FillGap(space - len(window)))
	builder.WriteString(footer)
	return builder.String()
}

// ---- Local methods

func (model IndexModel) getSpace() (footer string, space int) {
	var f string

	for _, notification := range model.notifications {
		f += model.readNotification(notification)
	}
	for _, notification := range model.Screens[model.Current].Notifications() {
		f += model.readNotification(notification)
	}

	f += model.helpWidget.View()
	return f, model.height - TITLE_HEIGHT - lipgloss.Height(f)
}

func (model IndexModel) limitScroll(space int) IndexModel {
	limit := screen.ContentHeight(model.Screens[model.Current], model.width-1) - space + 1

	if model.scroll >= limit {
		model.scroll = limit
	}

	if model.scroll < 0 {
		model.scroll = 0
	}

	return model
}

// ---- Utils

func (model IndexModel) readNotification(notification tools.Notification) string {
	if len(notification.Text) == 0 {
		return ""
	}

	switch notification.Type {
	case tools.HELP:
		return style.HelpNotification.RenderLine(model.width, notification.Text)
	case tools.TASK:
		return style.TaskNotification.RenderLine(model.width, notification.Text)
	case tools.ERROR:
		return style.ErrorNotification.RenderLine(model.width, notification.Text)
	}

	return notification.Text
}

func setupScreen(s screen.ScreenTag) screen.Screen {
	switch s {
	case screen.USER_SCREEN:
		return screen.NewUserScreen()
	}

	return screen.NewHomeScreen()
}

// ---- Messages

type ErrTaskMsg struct {
	Task string
	Err  error
}

type EndTaskMsg struct {
	Task string
}

func loadConfig() tea.Msg {
	err := config.MCVMConfig.DecodeFromFile(mcvm.ConfigFile("mcvm.json"))
	if err != nil {
		return ErrTaskMsg{"load+mcvm.json", err}
	}

	return EndTaskMsg{Task: "load+mcvm.json"}
}

func saveConfig() tea.Msg {
	err := config.MCVMConfig.EncodeToFile(mcvm.ConfigFile("mcvm.json"))
	if err != nil {
		return ErrTaskMsg{"save+mcvm.json", err}
	}

	return EndTaskMsg{Task: "save+mcvm.json"}
}

func loadAuth() tea.Msg {
	err := auth.MCVMAuth.DecodeFromFile(mcvm.DataFile("internal", "auth", "db.json"))
	if err != nil {
		return ErrTaskMsg{"load+db.json", err}
	}

	return EndTaskMsg{Task: "load+db.json"}
}
