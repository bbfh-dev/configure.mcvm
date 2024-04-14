package styles

import (
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/utils"
	"github.com/charmbracelet/lipgloss"
)

var (
	text  = lipgloss.NewStyle()
	title = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true, true, true, true).
		Foreground(ForegroundColor)
	infoBody = lipgloss.NewStyle().
			Border(lipgloss.ThickBorder(), false, false, false, true).
			BorderForeground(lipgloss.Color(ACCENT_COLOR)).
			PaddingLeft(1)
	infoTitle = infoBody.Copy().Foreground(lipgloss.Color(ACCENT_COLOR))
)

func Title(width int, contents ...string) string {
	return title.Width(width-2).Align(lipgloss.Center).Render(contents...) + "\n"
}

func Info(width int, contents ...string) string {
	return strings.Join([]string{
		infoTitle.Width(width).Render(utils.WithIcon(utils.INFO_ICON, "INFO:")),
		infoBody.Width(width).Render(strings.Join(contents, "\n")),
	}, "\n") + "\n"
}

func Center(width int, contents ...string) string {
	return text.Width(width).Align(lipgloss.Center).Render(contents...)
}
