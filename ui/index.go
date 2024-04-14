package ui

import (
	"github.com/bbfh-dev/configure.mcvm/ui/layout"
	"github.com/bbfh-dev/configure.mcvm/ui/styles"
	"github.com/bbfh-dev/configure.mcvm/ui/utils"
	"github.com/bbfh-dev/configure.mcvm/ui/widgets"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

const DEFAULT_SPACING = 1

const WELCOME_MESSAGE = "Welcome! You can see available actions on the bottom of the screen, press `?` to toggle the help message."

type IndexModel struct {
	Title     string
	Page      int
	Width     int
	Height    int
	screen    *layout.Screen
	helpModel tea.Model
}

func NewIndexModel() *IndexModel {
	return &IndexModel{
		Title: utils.WithIcon(utils.GEAR_ICON, "MCVM Configuration"),
		Page:  0,
		helpModel: widgets.NewHelpWidget(
			utils.ExtendWithDefaultKeybinds(utils.Keybinds{
				"+up": key.NewBinding(
					key.WithKeys("up", "k"),
					key.WithHelp("↑/k", "up"),
				),
				"+down": key.NewBinding(
					key.WithKeys("down", "j"),
					key.WithHelp("↓/j", "down"),
				),
				"+left": key.NewBinding(
					key.WithKeys("left", "h", "pgup"),
					key.WithHelp("←/l/pgup", "previous page"),
				),
				"+right": key.NewBinding(
					key.WithKeys("right", "l", "pgdown"),
					key.WithHelp("→/l/pgdn", "next page"),
				),
			}),
		),
	}
}

func (model *IndexModel) Init() tea.Cmd {
	model.screen = &layout.Screen{
		Title:     model.Title,
		PageIndex: &model.Page,
		Spacing:   DEFAULT_SPACING,
		Widgets: []layout.Widget{
			{TextStyle: styles.Info, Text: WELCOME_MESSAGE},
		},
		StatusWidgets: []layout.Widget{{Model: &model.helpModel}},
	}

	return nil
}

func (model *IndexModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		model.Width = msg.Width
		model.Height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "h", "left", "pgup":
			if model.Page > 0 {
				model.Page -= 1
			}
			return model, nil
		case "l", "right", "pgdown":
			model.Page += 1
			return model, nil
		case "ctrl+c", "q":
			return model, tea.Quit
		}
	}

	var cmd tea.Cmd
	model.helpModel, cmd = model.helpModel.Update(message)

	return model, tea.Batch(cmd)
}

func (model *IndexModel) View() string {
	if model.screen == nil {
		model.Init()
	}

	return model.screen.Render(
		model.Width,
		model.Height,
	)
}
