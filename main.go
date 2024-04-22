package main

import (
	"fmt"
	"os"

	"github.com/bbfh-dev/configure.mcvm/cli"
	"github.com/bbfh-dev/configure.mcvm/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func init() {
	cli.ParseFlags()
}

func main() {
	defer cli.SetupLogger().Close()

	if _, err := tea.NewProgram(ui.NewIndexModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}
}
