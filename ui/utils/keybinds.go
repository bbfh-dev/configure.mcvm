package utils

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
		key.WithKeys("q", "esc", "ctrl+c"),
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

	// var index int
	// var i int
	// for _, value := range keybinds {
	// 	keys[index] = append(keys[index], value)
	// 	if i++; i > 2 {
	// 		keys = append(keys, []key.Binding{})
	// 		index += 1
	// 		i = 0
	// 	}
	// }

	return keys

	// var keys []key.Binding
	// for key, value := range keybinds {
	// 	if key == "help" || key == "quit" {
	// 		continue
	// 	}
	// 	keys = append(keys, value)
	// }
	//
	// return [][]key.Binding{
	// 	keys,
	// 	{keybinds["help"], keybinds["quit"]},
	// }
}
