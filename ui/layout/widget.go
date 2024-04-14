package layout

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Widget struct {
	Header    bool
	Model     *tea.Model
	TextStyle func(int, ...string) string
	Text      string
}

func (widget *Widget) isModel() bool {
	return len(widget.Text) == 0
}

func (widget *Widget) Render(width int) string {
	if widget.isModel() {
		model := *widget.Model
		return model.View()
	}

	return widget.TextStyle(width, widget.Text)
}
