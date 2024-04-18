package styles

import "github.com/charmbracelet/lipgloss"

var withBorder = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true)
var quote = lipgloss.NewStyle().BorderLeft(true).BorderStyle(lipgloss.ThickBorder()).PaddingLeft(1)

var (
	Text   = defaultRenderer{lipgloss.NewStyle()}
	Bright = defaultRenderer{lipgloss.NewStyle().Foreground(ForegroundColor)}
	Header = defaultRenderer{lipgloss.NewStyle().PaddingLeft(2).Foreground(AccentColor)}
	Title  = defaultRenderer{
		withBorder.Copy().Align(lipgloss.Center).Foreground(ForegroundColor),
	}
	Hint         = defaultRenderer{lipgloss.NewStyle().Foreground(HintColor)}
	ListItem     = defaultRenderer{lipgloss.NewStyle().Foreground(InactiveColor)}
	CursorItem   = defaultRenderer{lipgloss.NewStyle().Foreground(ForegroundColor)}
	Notification = defaultRenderer{lipgloss.NewStyle().Foreground(NotificationColor)}
	Info         = quoteRenderer{quote.Copy().BorderForeground(AccentColor).Foreground(AccentColor)}
)
