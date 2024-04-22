package widget

import (
	"fmt"
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/style"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/bbfh-dev/configure.mcvm/ui/widget/field"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type FormWidget struct {
	focused bool
	Editing bool
	Width   int
	Cursor  int
	Keys    tools.Keybinds
	Fields  []field.Field

	Messages     []tools.Message
	Notification tools.Notification
}

func NewFormWidget(fields []field.Field, focused bool) FormWidget {
	return FormWidget{
		Fields: fields,
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
			"item.exit": key.NewBinding(
				key.WithKeys("esc"),
				key.WithHelp("esc", "exit editting"),
			),
		},
		focused: true,
	}
}

func (widget FormWidget) Init() tea.Cmd {
	return nil
}

func (widget FormWidget) Update(raw tea.Msg) (tea.Model, tea.Cmd) {
	if !widget.focused {
		return widget, nil
	}

	switch msg := raw.(type) {
	case tea.KeyMsg:
		if widget.Editing {
			switch {
			case key.Matches(msg, widget.Keys["item.exit"]):
				widget.Editing = false
			default:
				widget.Fields[widget.Cursor].Type = widget.Fields[widget.Cursor].Type.Update(msg)
			}
			break
		}
		switch {

		case key.Matches(msg, widget.Keys["move.up"]):
			if widget.Cursor > 0 {
				widget.Cursor -= 1
			} else {
				widget.Messages = append(widget.Messages, tools.OverflowTopMsg{})
			}

		case key.Matches(msg, widget.Keys["move.down"]):
			if widget.Cursor < len(widget.Fields)-1 {
				widget.Cursor += 1
			} else {
				widget.Messages = append(widget.Messages, tools.OverflowBottomMsg{})
			}

		case key.Matches(msg, widget.Keys["item.select"]):
			widget.Editing = true
		}
	}

	widget.Notification = tools.Notification{
		Type: tools.HELP,
		Text: widget.Fields[widget.Cursor].Description,
	}

	return widget, nil
}

func (widget FormWidget) View() string {
	var builder strings.Builder

	for i, field := range widget.Fields {
		name := fmt.Sprintf(" %s: ", field.Name)
		if widget.focused && i == widget.Cursor {
			builder.WriteString(
				style.Bright.Render(
					widget.Width-2,
					name,
				),
			)
			builder.WriteString(
				style.Field.Render(
					widget.Width-2-len(name),
					fmt.Sprintf("[ %s ]", field.Type.String()),
				) + "\n",
			)

			if widget.Editing {
				builder.WriteString(field.Type.View(widget.Width-2) + "\n")
			}
		} else {
			builder.WriteString(
				style.Inactive.Render(widget.Width-2, fmt.Sprintf("  %s: [%s]", field.Name, field.Type.String())) + "\n",
			)
		}
	}

	return builder.String()
}

func (widget FormWidget) Clear() FormWidget {
	widget.Messages = make([]tools.Message, 0)
	widget.Notification = tools.Notification{}

	return widget
}

func (widget FormWidget) Blur() FormWidget {
	widget.focused = false
	return widget
}

func (widget FormWidget) Focus() FormWidget {
	widget.focused = true
	return widget
}

// ---- Messages

type SaveMsg struct{}
