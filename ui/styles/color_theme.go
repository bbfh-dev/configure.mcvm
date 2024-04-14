package styles

import "github.com/charmbracelet/lipgloss"

const (
	ACCENT_COLOR = "#328bff"
	BLACK_COLOR  = "#000000"
	WHITE_COLOR  = "#ffffff"
)

var (
	ForegroundColor = lipgloss.AdaptiveColor{Light: BLACK_COLOR, Dark: WHITE_COLOR}
)
