package main

import (
	"fmt"
	"os"

	"github.com/bbfh-dev/configure.mcvm/cli"
	"github.com/bbfh-dev/configure.mcvm/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jessevdk/go-flags"
)

func init() {
	_, err := flags.Parse(&cli.Flags)
	if err != nil {
		os.Exit(0)
	}
}

func main() {
	defer cli.SetupLogger().Close()
	cli.Logger("Started the program").Info()

	program := tea.NewProgram(ui.MainModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}
}
