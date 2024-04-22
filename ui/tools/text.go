package tools

import (
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var dft = lipgloss.NewStyle()

// Fills vertical gap with new lines
func FillGap(height int) string {
	return strings.Repeat("\n", max(0, height))
}

// Calculates the height and position of a scrollbar
func ScrollBar(windowHeight int, contentHeight int, scroll int) (height int, position int) {
	ratio := float64(windowHeight) / float64(contentHeight)
	height = int(math.Max(float64(windowHeight)*ratio, 1))
	position = int(math.Round(float64(scroll) * ratio))

	if position == 0 && scroll != 0 {
		position = 1
	}

	if scroll+windowHeight >= contentHeight {
		position = windowHeight - height
	}

	return height, position
}

// Appends a text render with specified width
func AppendContent(contents *[]string, text string, width int) {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = dft.Width(width).Render(line)
	}
	*contents = append(*contents, lines...)
}
