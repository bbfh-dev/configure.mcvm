package widgets

import (
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type HelpWidget struct {
	keys      utils.Keybinds
	helpModel help.Model
}

func NewHelpWidget(keys utils.Keybinds) HelpWidget {
	return HelpWidget{
		keys:      keys,
		helpModel: help.New(),
	}
}

func (model HelpWidget) Init() tea.Cmd {
	return nil
}

func (model HelpWidget) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		model.helpModel.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, model.keys["help"]):
			model.helpModel.ShowAll = !model.helpModel.ShowAll
		}
	}

	return model, nil
}

func (model HelpWidget) View() string {
	var contents string
	helpView := model.helpModel.View(model.keys)
	height := 1 - strings.Count(contents, "\n") - strings.Count(helpView, "\n")

	return "\n" + contents + strings.Repeat("\n", max(0, height)) + helpView
}
