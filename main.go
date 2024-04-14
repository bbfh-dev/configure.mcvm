package main

import (
	"fmt"
	"os"

	"github.com/bbfh-dev/configure.mcvm/cli"
	"github.com/bbfh-dev/configure.mcvm/mcvm"
	"github.com/bbfh-dev/configure.mcvm/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jessevdk/go-flags"
)

func main() {
	defer cli.SetupLogger().Close()

	_, err := flags.Parse(&cli.Flags)
	if err != nil {
		os.Exit(0)
	}

	err = mcvm.MCVMConfig.DecodeFromFile()
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		os.Exit(1)
	}

	// err := mcvm.MCVMConfig.DecodeFromFile()

	// err := mcvm.DecodeFromFile()
	// if err != nil {
	// 	fmt.Printf("ERROR: %s", err)
	// 	os.Exit(1)
	// }
	//
	// //region: Watcher
	// watcher := cli.SetupWatcher(func(filename string) {
	// 	switch filename {
	// 	case "mcvm.json":
	// 		mcvm.DecodeFromFile()
	// 	}
	// })
	// defer watcher.Close()
	//
	// err = watcher.Add(mcvm.GetMCVMFile("mcvm.json"))
	// if err != nil {
	// 	cli.Logger(err.Error()).Error()
	// }
	// //endregion

	program := tea.NewProgram(ui.NewIndexModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}
}
