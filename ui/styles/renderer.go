package styles

import (
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/charmbracelet/lipgloss"
)

type Renderer interface {
	Render(int, ...string) string
}

type defaultRenderer struct {
	style lipgloss.Style
}

func (renderer *defaultRenderer) Render(width int, lines ...string) string {
	content := strings.Join(lines, "\n")
	return renderer.style.Width(width - 2).Render(content)
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
