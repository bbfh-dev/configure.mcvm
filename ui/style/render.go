package style

import (
	"fmt"
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/charmbracelet/lipgloss"
)

type Renderer interface {
	Render(int, ...string) string
	RenderLine(int, ...string) string
}

type literalRenderer struct {
	style lipgloss.Style
}

func (renderer *literalRenderer) Render(width int, lines ...string) string {
	content := strings.Join(lines, "\n")
	return renderer.style.Width(width).Render(content)
}

func (renderer *literalRenderer) RenderLine(width int, lines ...string) string {
	return renderer.Render(width, lines...) + "\n"
}

type quoteRenderer struct {
	style lipgloss.Style
}

func (renderer *quoteRenderer) Render(width int, lines ...string) string {
	contents := renderer.style.Width(width).Render(tools.WithIcon(tools.INFO_ICON, "INFO"))
	for _, line := range lines {
		contents += "\n" + quote.BorderForeground(AccentColor).Width(width).Render(line)
	}

	return contents
}

func (renderer *quoteRenderer) RenderLine(width int, lines ...string) string {
	return renderer.Render(width, lines...) + "\n"
}

type notificationRenderer struct {
	icon  string
	style lipgloss.Style
}

func (renderer *notificationRenderer) Render(width int, lines ...string) string {
	return renderer.style.Width(width).
		Render(fmt.Sprintf(
			"%s %s",
			renderer.icon,
			strings.Join(lines, "\n"),
		))
}

func (renderer *notificationRenderer) RenderLine(width int, lines ...string) string {
	return renderer.Render(width, lines...) + "\n"
}
