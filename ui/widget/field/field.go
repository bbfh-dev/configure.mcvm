package field

import tea "github.com/charmbracelet/bubbletea"

type Field struct {
	Name        string
	Description string
	Type        FieldType
}

func NewField(name string, description string, field FieldType) Field {
	return Field{
		Name:        name,
		Description: description,
		Type:        field,
	}
}

type FieldType interface {
	Id() string
	String() string
	Update(key tea.KeyMsg) FieldType
	View(width int) string
	Set(any) FieldType
}
