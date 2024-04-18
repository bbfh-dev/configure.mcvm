package screens

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbfh-dev/configure.mcvm/cli"
	"github.com/bbfh-dev/configure.mcvm/ui/styles"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Hardcoded since headers are only used for titles CURRENTLY.
const HEADER_HEIGHT = 3
const SPACING = 2

type Screen interface {
	tea.Model
	SetWidth(int) Screen
	Height() int
	SetHeight(int) Screen
	Scroll() int
	MaxScroll() int
	SetScroll(int) Screen
	SetMaxScroll(int) Screen
	Render() (string, []string, *int)
	Keybinds() tools.Keybinds
	DidExit() bool
	Exit() bool
	Reload() Screen
	Notifications() []string
	RenderContents() []string
}

func RenderScreen(
	scrollbar Scrollbar,
	width *int,
	height *int,
	header string,
	footer string,
	lines []string,
) string {
	contentHeight := *height - HEADER_HEIGHT - lipgloss.Height(footer) - SPACING
	var content string
	var margin string

	offset := min(len(lines)-1, *scrollbar.scroll)
	index := 0

	for i, line := range lines[offset:] {
		if i > contentHeight {
			break
		}

		margin = strings.Repeat(" ", max(0, *width-lipgloss.Width(line)-1))
		content += line + margin + scrollbar.getScrollbarAt(index) + "\n"
		index += 1
	}

	padding := strings.Repeat("\n", max(1, contentHeight-len(lines)+2))
	return header + "\n" + content + padding + footer
}

type Scrollbar struct {
	entries []bool
	scroll  *int
}

func (scrollbar Scrollbar) getScrollbarAt(position int) string {
	if scrollbar.entries[position] {
		return styles.Bright.Render(1, "┃")
	}

	return styles.Hint.Render(1, "┃")
}

func CalculateScrollbar(height int, contentHeight int, scroll *int) Scrollbar {
	if contentHeight <= height {
		var bar = make([]bool, height)
		for i := range height {
			bar[i] = false
		}

		return Scrollbar{entries: bar, scroll: scroll}
	}

	content := float64(contentHeight)
	percentage := float64(height) / content
	barHeight := math.Floor(float64(height) * percentage)

	offset := int((float64(height) - float64(barHeight)) * (float64(*scroll) / content))
	barLength := min(int(barHeight)+offset, height-2)
	cli.Logger(fmt.Sprintf("%v = %v", barLength, height-1)).Debug()

	var bar = make([]bool, height)
	for i := range height {
		if i < offset {
			bar[i] = false
			continue
		}
		bar[i] = i <= barLength
		cli.Logger(fmt.Sprintf("%v - %v", i, i <= barLength+1)).Debug()
	}

	return Scrollbar{entries: bar, scroll: scroll}
}
