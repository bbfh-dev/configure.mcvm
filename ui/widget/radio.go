package widget

import (
	"fmt"
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/component"
	"github.com/bbfh-dev/configure.mcvm/ui/style"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type RadioWidget struct {
	focused bool
	Width   int
	Cursor  int
	List    component.List
	Keys    tools.Keybinds

	Messages     []tools.Message
	Notification tools.Notification
}

func NewRadioWidget(list component.List, focused bool) RadioWidget {
	return RadioWidget{
		List: list,
		Keys: tools.Keybinds{
			"move.up": key.NewBinding(
				key.WithKeys("shift+up", "k"),
				key.WithHelp("󰘶 ↑/k", "up"),
			),
			"move.down": key.NewBinding(
				key.WithKeys("shift+down", "j"),
				key.WithHelp("󰘶 ↓/j", "down"),
			),
			"item.select": key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("󰌑", "select"),
			),
		},
		focused: focused,
	}
}

func (widget RadioWidget) Init() tea.Cmd {
	return nil
}

func (widget RadioWidget) Update(raw tea.Msg) (tea.Model, tea.Cmd) {
	if !widget.focused {
		return widget, nil
	}

	switch msg := raw.(type) {
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, widget.Keys["move.up"]):
			if widget.Cursor > 0 {
				widget.Cursor -= 1
			} else {
				widget.Messages = append(widget.Messages, tools.OverflowTopMsg{})
			}

		case key.Matches(msg, widget.Keys["move.down"]):
			if widget.Cursor < widget.List.Size()-1 {
				widget.Cursor += 1
			} else {
				widget.Messages = append(widget.Messages, tools.OverflowBottomMsg{})
			}

		case key.Matches(msg, widget.Keys["item.select"]):
			widget.Messages = append(widget.Messages, component.SelectItemMsg{Item: widget.List.Items()[widget.Cursor].Id()})
		}
	}

	widget.Notification = tools.Notification{
		Type: tools.HELP,
		Text: widget.List.Items()[widget.Cursor].Description(),
	}

	return widget, nil
}

func (widget RadioWidget) View() string {
	var builder strings.Builder

	for i, item := range widget.List.Items() {
		if widget.focused && i == widget.Cursor {
			builder.WriteString(
				style.Bright.Render(widget.Width-2, fmt.Sprintf(" %s", item.Title())) + "\n",
			)
		} else {
			builder.WriteString(
				style.Inactive.Render(widget.Width-2, fmt.Sprintf(" %s", item.Title())) + "\n",
			)
		}
	}

	return builder.String()
}

func (widget RadioWidget) Clear() RadioWidget {
	widget.Messages = make([]tools.Message, 0)
	widget.Notification = tools.Notification{}

	return widget
}

func (widget RadioWidget) Focus() RadioWidget {
	widget.focused = true
	return widget
}

func (widget RadioWidget) Blur() RadioWidget {
	widget.focused = false
	return widget.Clear()
}
