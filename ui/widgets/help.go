package widgets

import (
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type HelpWidget struct {
	Id          string
	DefaultKeys []string
	Keys        tools.Keybinds
	helpModel   help.Model
}

func NewHelpWidget(keys tools.Keybinds) HelpWidget {
	helpModel := help.New()
	helpModel.ShowAll = true

	var keyList []string
	for key := range keys {
		keyList = append(keyList, key)
	}

	return HelpWidget{
		Id:          "default",
		DefaultKeys: keyList,
		Keys:        keys,
		helpModel:   helpModel,
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
		case key.Matches(msg, model.Keys["help"]):
			model.helpModel.ShowAll = !model.helpModel.ShowAll
		}
	}

	return model, nil
}

func (model HelpWidget) View() string {
	helpView := model.helpModel.View(model.Keys)
	height := 1 - strings.Count(helpView, "\n")

	return strings.Repeat("\n", max(0, height)) + helpView
}
