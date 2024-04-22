package cli

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Log(format string, v ...any) {
	log.Printf(format, v...)
}

func SetupLogger() *os.File {
	file, err := tea.LogToFile(Flags.Log, "")
	if err != nil {
		fmt.Println("FATAL: ", err)
		os.Exit(1)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime)

	return file
}
