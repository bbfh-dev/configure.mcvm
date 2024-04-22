package component

// An available UI option.
//
// Use it as a slice, first element should be the default option.
type Choice struct {
	id      string
	name    string
	docs    string
	Allowed bool
}

func (item Choice) Id() string {
	return item.id
}

func (item Choice) Title() string {
	return item.name
}

func (item Choice) Description() string {
	return item.docs
}

type ChoiceList struct {
	items *[]Choice
}

func NewChoiceList(items *[]Choice) ChoiceList {
	return ChoiceList{items: items}
}

func (list ChoiceList) Items() []Item {
	var items = make([]Item, list.Size())

	for i, item := range *list.items {
		items[i] = item
	}

	return items
}

func (list ChoiceList) Size() int {
	return len(*list.items)
}

var UserTypeOptions []Choice = []Choice{
	{
		id:      "microsoft",
		name:    "Microsoft",
		docs:    "A normal Minecraft account",
		Allowed: true,
	},
	{
		id:      "demo",
		name:    "Demo",
		docs:    "An account that owns a demo of the game",
		Allowed: true,
	},
	{
		id:      "unverified",
		name:    "Unverified",
		docs:    "An unverified or 'cracked' account",
		Allowed: true,
	},
}

var InstanceTypeOptions []Choice = []Choice{
	{
		id:      "client",
		name:    "Client",
		docs:    "The client side",
		Allowed: true,
	},
	{
		id:      "server",
		name:    "server",
		docs:    "The server side",
		Allowed: true,
	},
}

var PackageStabilityOptions []Choice = []Choice{
	{
		id:      "stable",
		name:    "Stable",
		docs:    "Use the most recent stable version",
		Allowed: true,
	},
	{
		id:      "latest",
		name:    "Latest",
		docs:    "Use most recent available version",
		Allowed: true,
	},
}

var JavaOptions []Choice = []Choice{
	{
		id:      "auto",
		name:    "Auto",
		docs:    "Automatically picks or downloads the best Java flavor for your system",
		Allowed: true,
	},
	{
		id:      "system",
		name:    "System",
		docs:    "It will try to find an existing installation on your system",
		Allowed: true,
	},
	{
		id:      "adoptium",
		name:    "Adoptium",
		docs:    "Adoptium Installation",
		Allowed: true,
	},
	{
		id:      "zulu",
		name:    "Zulu",
		docs:    "Zulu Installation",
		Allowed: true,
	},
	{
		id:      "graalvm",
		name:    "GraalVM",
		docs:    "GraalVM Installation",
		Allowed: true,
	},
}

var PresetOptions []Choice = []Choice{
	{
		id:      "none",
		name:    "None",
		docs:    "No changes will be applied",
		Allowed: true,
	},
	{
		id:      "aikars",
		name:    "Aikars",
		docs:    "A popular set of tuned arguments for better performance. This works better for servers that have a lot of available memory (8GB+) and is not recommended otherwise",
		Allowed: true,
	},
	{
		id:      "krusic",
		name:    "Krusic",
		docs:    "Another set of performance arguments",
		Allowed: true,
	},
	{
		id:      "obydux",
		name:    "Obydux",
		docs:    "Another set of performance arguments",
		Allowed: true,
	},
}

var SnapshotsStorageOptions []Choice = []Choice{
	{
		id:      "archive",
		name:    "Archive",
		docs:    "Store snapshots as an archive",
		Allowed: true,
	},
	{
		id:      "folder",
		name:    "Folder",
		docs:    "Store snapshots as a folder",
		Allowed: true,
	},
}

var PackagePermissionOptions []Choice = []Choice{
	{
		id:      "standard",
		name:    "Standard",
		docs:    "Default package permissions",
		Allowed: true,
	},
	{
		id:      "restricted",
		name:    "Restricted",
		docs:    "Packages you do not trust",
		Allowed: true,
	},
	{
		id:      "elevated",
		name:    "Elevated",
		docs:    "Packages that you trust and want to provide access to special commands for",
		Allowed: true,
	},
}

var CachingStrategyOptions []Choice = []Choice{
	{
		id:      "all",
		name:    "All",
		docs:    "Cache all packages whenever you run the `package sync` command",
		Allowed: true,
	},
	{
		id:      "none",
		name:    "None",
		docs:    "Never cache any scripts",
		Allowed: true,
	},
	{
		id:      "lazy",
		name:    "Lazy",
		docs:    "Cache only when a package is requested",
		Allowed: true,
	},
}
