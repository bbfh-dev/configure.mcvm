package cli

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

type Logger string

func (logger Logger) log(prefix string) {
	log.Printf("%s: %s", prefix, logger)
}

func (logger Logger) Info() {
	logger.log("INFO")
}

func (logger Logger) Debug() {
	logger.log("DEBUG")
}

func (logger Logger) Error() {
	logger.log("ERROR")
}

func SetupLogger() *os.File {
	if len(Flags.Log) <= 0 {
		Flags.Log = filepath.Join(os.TempDir(), "mcvmconf.log")
	}

	file, err := tea.LogToFile(Flags.Log, "")
	if err != nil {
		fmt.Println("FATAL: ", err)
		os.Exit(1)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime)

	return file
}
