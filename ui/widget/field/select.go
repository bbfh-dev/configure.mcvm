package field

import (
	"fmt"
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/component"
	"github.com/bbfh-dev/configure.mcvm/ui/style"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	tea "github.com/charmbracelet/bubbletea"
)

type SelectField struct {
	cursor int
	items  []component.Item
	keys   tools.Keybinds
}

func NewSelectField(list component.List) SelectField {
	return SelectField{items: list.Items()}
}

func (field SelectField) Id() string {
	return field.items[field.cursor].Id()
}

func (field SelectField) String() string {
	return field.items[field.cursor].Title()
}

func (field SelectField) Update(bind tea.KeyMsg) FieldType {
	switch bind.String() {

	case "k", "shift+up":
		if field.cursor > 0 {
			field.cursor -= 1
		}

	case "j", "shift+down":
		if field.cursor < len(field.items)-1 {
			field.cursor += 1
		}
	}
	return field
}

func (field SelectField) View(width int) string {
	var builder strings.Builder

	for i, item := range field.items {
		if i == field.cursor {
			builder.WriteString(style.Bright.Render(width, fmt.Sprintf(" %s ", item.Title())))
			builder.WriteString(style.Details.RenderLine(width, item.Description()))
		} else {
			builder.WriteString(
				style.Inactive.Render(width, fmt.Sprintf(" %s", item.Title())) + "\n",
			)
		}
	}

	return builder.String()
}

func (field SelectField) Set(value any) FieldType {
	return field
}
