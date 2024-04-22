package tools

import (
	tea "github.com/charmbracelet/bubbletea"
)

// A util function that calls the model.Update() method,
// appends the command, and maps the return value from tea.Model
func UpdateModel[T tea.Model](commands *[]tea.Cmd, model tea.Model, raw tea.Msg) T {
	m, cmd := model.Update(raw)
	*commands = append(*commands, cmd)

	return m.(T)
}
