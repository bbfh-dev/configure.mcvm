package widgets

import (
	"fmt"

	"github.com/bbfh-dev/configure.mcvm/ui/styles"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type RadioWidget struct {
	focused      bool
	Title        string
	Width        int
	Height       int
	List         tools.List
	Cursor       int
	Selection    string
	Keys         tools.Keybinds
	Notification string // Shows up in the footer
}

func NewRadioWidget(title string, list tools.List) RadioWidget {
	return RadioWidget{
		Title: title,
		List:  list,
		Keys: tools.Keybinds{
			"10+previous": key.NewBinding(
				key.WithKeys("left", "k"),
				key.WithHelp("←/k", "previous"),
			),
			"10+next": key.NewBinding(
				key.WithKeys("right", "j"),
				key.WithHelp("→/j", "next"),
			),
			"11+select": key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("󰌑", "select"),
			),
		},
	}
}

func (widget RadioWidget) Init() tea.Cmd {
	return nil
}

func (widget RadioWidget) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if !widget.focused {
		widget.Notification = ""
		return widget, nil
	}

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, widget.Keys["10+previous"]):
			widget.Selection = ""
			if widget.Cursor > 0 {
				widget.Cursor -= 1
			}
		case key.Matches(msg, widget.Keys["10+next"]):
			widget.Selection = ""
			if widget.Cursor < len(widget.List.Items())-1 {
				widget.Cursor += 1
			}
		case key.Matches(msg, widget.Keys["11+select"]):
			widget.Selection = widget.List.Items()[widget.Cursor].Id()
		}
	}

	widget.Notification = widget.List.Items()[widget.Cursor].Description()

	return widget, nil
}

func (widget RadioWidget) View() string {
	content := styles.Header.Render(widget.Width, widget.Title) + "\n"
	for i, item := range widget.List.Items() {
		text := fmt.Sprintf("%s %s", tools.IconFallback(tools.RIGHT_ARROW_ICON, ">"), item.Name())

		if i == widget.Cursor {
			content += styles.CursorItem.Render(widget.Width, text)
		} else {
			content += styles.ListItem.Render(widget.Width, text)
		}

		content += "\n"
	}
	return content
}

func (widget RadioWidget) Focus() RadioWidget {
	widget.focused = true

	return widget
}

func (widget RadioWidget) Blur() RadioWidget {
	widget.focused = false

	return widget
}

func (widget RadioWidget) HasSelected() bool {
	return len(widget.Selection) != 0
}

type UpdateListMsg struct{}
