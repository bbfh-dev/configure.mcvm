package layout

import (
	"strings"

	"github.com/bbfh-dev/configure.mcvm/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

type Screen struct {
	Title         string
	PageIndex     *int
	Spacing       int
	Widgets       []Widget
	StatusWidgets []Widget
}

func (screen Screen) Render(width int, height int) string {
	if width == 0 || height == 0 {
		return "Not enough screen space!"
	}

	pages := []string{styles.Title(width, screen.Title)}
	index := 0
	availableHeight := height - lipgloss.Height(pages[0])

	var footer string
	for _, widget := range screen.StatusWidgets {
		footer += widget.Render(width)
	}
	footerHeight := lipgloss.Height(footer) + 3

	for _, widget := range screen.Widgets {
		text := widget.Render(width)

		if availableHeight < footerHeight {
			pages = append(pages, styles.Title(width, screen.Title))
			index += 1
			availableHeight = height - lipgloss.Height(pages[index]) - 1
		} else {
			availableHeight -= lipgloss.Height(text) + screen.Spacing
		}

		pages[index] += strings.Repeat("\n", screen.Spacing)
		pages[index] += text
	}

	if len(pages) == 1 {
		*screen.PageIndex = 0
		return pages[0] + strings.Repeat(
			"\n",
			max(0, height-lipgloss.Height(pages[0])-lipgloss.Height(footer)),
		) + footer
	}

	*screen.PageIndex = min(*screen.PageIndex, index)
	viewport := pages[*screen.PageIndex] + strings.Repeat(
		"\n",
		max(0, height-lipgloss.Height(pages[*screen.PageIndex])-lipgloss.Height(footer)),
	)

	pager := make([]string, index+1)
	for i := range index + 1 {
		if i == *screen.PageIndex {
			pager[i] = "●"
		} else {
			pager[i] = "○"
		}
	}

	viewport += styles.Center(width, strings.Join(pager, " ")) + footer
	return viewport + strings.Repeat("\n", max(0, height-lipgloss.Height(viewport)))
}
