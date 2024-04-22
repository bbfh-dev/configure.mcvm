package tools

import (
	"sort"

	"github.com/charmbracelet/bubbles/key"
)

const FULL_HELP_HEIGHT = 4

type Keybinds map[string]key.Binding

var HelpKeybind = key.NewBinding(
	key.WithKeys("?"),
	key.WithHelp("?", "help"),
)

var QuitKeybind = key.NewBinding(
	key.WithKeys("q", "ctrl+c"),
	key.WithHelp("q", "quit"),
)

var GlobalKeybinds = Keybinds{
	"scroll.up": key.NewBinding(
		key.WithKeys("K", "up"),
		key.WithHelp("󰘶 K/󱕑", "scroll up"),
	),
	"scroll.down": key.NewBinding(
		key.WithKeys("J", "down"),
		key.WithHelp("󰘶 J/󱕐", "scroll down"),
	),
	"goto.top": key.NewBinding(
		key.WithKeys("ctrl+u", "home"),
		key.WithHelp("^U/home", "go to top"),
	),
	"goto.bottom": key.NewBinding(
		key.WithKeys("ctrl+d", "end"),
		key.WithHelp("^D/end", "go to bottom"),
	),
}

func (keybinds Keybinds) ShortHelp() []key.Binding {
	return []key.Binding{HelpKeybind, QuitKeybind}
}

func (keybinds Keybinds) FullHelp() [][]key.Binding {
	keys := [][]key.Binding{{}}

	var binds []string
	for key := range keybinds {
		binds = append(binds, key)
	}
	sort.Strings(binds)

	i := 0
	for _, k := range binds {
		keys[len(keys)-1] = append(keys[len(keys)-1], keybinds[k])
		if i++; i%FULL_HELP_HEIGHT == 0 {
			keys = append(keys, []key.Binding{})
		}
	}

	return keys
}
