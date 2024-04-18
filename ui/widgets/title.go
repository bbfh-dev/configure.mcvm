package widgets

import (
	"github.com/bbfh-dev/configure.mcvm/ui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type TitleWidget struct {
	Title string
	Width int
}

func NewTitleWidget(title string) TitleWidget {
	return TitleWidget{
		Title: title,
	}
}

func (widget TitleWidget) Init() tea.Cmd {
	return nil
}

func (widget TitleWidget) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return widget, nil
}

func (widget TitleWidget) View() string {
	return styles.Title.Render(widget.Width, widget.Title)
}
