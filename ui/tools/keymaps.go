package tools

import (
	"sort"

	"github.com/charmbracelet/bubbles/key"
)

// Keybinds defines a set of keybindings.
type Keybinds map[string]key.Binding

func ExtendWithDefaultKeybinds(keybinds Keybinds) Keybinds {
	keybinds["help"] = key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	)
	keybinds["quit"] = key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	)
	return keybinds
}

// ShortHelp returns keybindings to be shown in the mini help view.
func (keybinds Keybinds) ShortHelp() []key.Binding {
	return []key.Binding{keybinds["help"], keybinds["quit"]}
}

// FullHelp returns keybindings for the expanded help view.
func (keybinds Keybinds) FullHelp() [][]key.Binding {
	keys := [][]key.Binding{{}}

	var mapKeys []string
	for key := range keybinds {
		mapKeys = append(mapKeys, key)
	}
	sort.Strings(mapKeys)

	i := 0
	for _, k := range mapKeys {
		keys[len(keys)-1] = append(keys[len(keys)-1], keybinds[k])
		if i++; i%2 == 0 {
			keys = append(keys, []key.Binding{})
		}
	}

	return keys
}
