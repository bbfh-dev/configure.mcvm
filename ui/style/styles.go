package style

import (
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/charmbracelet/lipgloss"
)

var quote = lipgloss.NewStyle().BorderLeft(true).BorderStyle(lipgloss.ThickBorder()).PaddingLeft(1)

var (
	Text     = literalRenderer{lipgloss.NewStyle().Foreground(RegularColor)}
	Bright   = literalRenderer{lipgloss.NewStyle().Foreground(ForegroundColor)}
	Inactive = literalRenderer{lipgloss.NewStyle().Foreground(InactiveColor)}
	Details  = literalRenderer{lipgloss.NewStyle().Foreground(TaskColor)}
	Field    = literalRenderer{lipgloss.NewStyle().Foreground(FieldColor)}
	Title    = literalRenderer{
		lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(InactiveColor).
			Align(lipgloss.Center).
			Foreground(ForegroundColor),
	}
	Info = quoteRenderer{quote.Copy().BorderForeground(AccentColor).Foreground(AccentColor)}
)
var (
	HelpNotification = notificationRenderer{
		string(tools.INFO_ICON),
		lipgloss.NewStyle().Foreground(RegularColor),
	}
	TaskNotification = notificationRenderer{
		string(tools.LOADING_ICON),
		lipgloss.NewStyle().Foreground(TaskColor),
	}
	ErrorNotification = notificationRenderer{
		string(tools.WARNING_ICON),
		lipgloss.NewStyle().Foreground(ErrorColor),
	}
)
