package widget

import (
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type HelpWidget struct {
	Keys      tools.Keybinds
	helpModel help.Model
}

func NewHelpWidget(keys tools.Keybinds) HelpWidget {
	keys["z.help"] = tools.HelpKeybind
	keys["z.quit"] = tools.QuitKeybind

	return HelpWidget{
		Keys:      keys,
		helpModel: help.New(),
	}
}

func (widget HelpWidget) Init() tea.Cmd {
	return nil
}

func (widget HelpWidget) Update(raw tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := raw.(type) {
	case tea.WindowSizeMsg:
		widget.helpModel.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, tools.HelpKeybind):
			widget.helpModel.ShowAll = !widget.helpModel.ShowAll
		}
	}

	return widget, nil
}

func (widget HelpWidget) View() string {
	return widget.helpModel.View(widget.Keys)
}
