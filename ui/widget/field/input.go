package field

import (
	"fmt"
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/style"
	tea "github.com/charmbracelet/bubbletea"
)

const MAX_STRING_LENGTH = 16

type InputField struct {
	text      string
	validator validator
}

func NewInputField(validator validator) InputField {
	return InputField{validator: validator}
}

func (field InputField) String() string {
	return fmt.Sprintf("%q", field.text)
}

func (field InputField) Id() string {
	return field.text
}

func (field InputField) Update(key tea.KeyMsg) FieldType {
	switch key.String() {
	case "backspace":
		if len(field.text) > 0 {
			field.text = field.text[:len(field.text)-1]
		}
	default:
		field.text += key.String()
	}

	field.text = field.validator.Validate(field.text)
	return field
}

func (field InputField) View(width int) string {
	return style.Bright.Render(width, fmt.Sprintf(" %s", field.text))
}

func (field InputField) Set(value any) FieldType {
	field.text = fmt.Sprintf("%s", value)
	return field
}

// ---- Validators

type validator interface {
	Validate(value string) string
}

type IdValidator struct{}

func (validator IdValidator) Validate(value string) string {
	result := strings.ReplaceAll(strings.ToLower(value), " ", "_")
	if len(result) > MAX_STRING_LENGTH {
		return result[:MAX_STRING_LENGTH]
	}

	return result
}
