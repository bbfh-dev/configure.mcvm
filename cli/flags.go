package cli

var Flags struct {
	Log   string `short:"l" long:"log" description:"Path to log file"`
	Icons bool   `short:"i" long:"icons" description:"Enable icons in TUI"`
}
