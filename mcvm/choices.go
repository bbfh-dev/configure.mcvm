package mcvm

// An available UI option.
// Use it as a slice, first element should be the default option.
type Choice struct {
	Id      string
	Name    string
	Docs    string
	Allowed bool
}

var UserTypeOptions []Choice = []Choice{
	{
		Id:      "microsoft",
		Name:    "Microsoft",
		Docs:    "A normal Minecraft account",
		Allowed: true,
	},
	{
		Id:      "demo",
		Name:    "Demo",
		Docs:    "An account that owns a demo of the game",
		Allowed: true,
	},
	{
		Id:      "unverified",
		Name:    "Unverified",
		Docs:    "An unverified or 'cracked' account",
		Allowed: true,
	},
}

var InstanceTypeOptions []Choice = []Choice{
	{
		Id:      "client",
		Name:    "Client",
		Docs:    "The client side",
		Allowed: true,
	},
	{
		Id:      "server",
		Name:    "server",
		Docs:    "The server side",
		Allowed: true,
	},
}

var PackageStabilityOptions []Choice = []Choice{
	{
		Id:      "stable",
		Name:    "Stable",
		Docs:    "Use the most recent stable version",
		Allowed: true,
	},
	{
		Id:      "latest",
		Name:    "Latest",
		Docs:    "Use most recent available version",
		Allowed: true,
	},
}

var JavaOptions []Choice = []Choice{
	{
		Id:      "auto",
		Name:    "Auto",
		Docs:    "Automatically picks or downloads the best Java flavor for your system",
		Allowed: true,
	},
	{
		Id:      "system",
		Name:    "System",
		Docs:    "It will try to find an existing installation on your system",
		Allowed: true,
	},
	{
		Id:      "adoptium",
		Name:    "Adoptium",
		Docs:    "Adoptium Installation",
		Allowed: true,
	},
	{
		Id:      "zulu",
		Name:    "Zulu",
		Docs:    "Zulu Installation",
		Allowed: true,
	},
	{
		Id:      "graalvm",
		Name:    "GraalVM",
		Docs:    "GraalVM Installation",
		Allowed: true,
	},
}

var PresetOptions []Choice = []Choice{
	{
		Id:      "none",
		Name:    "None",
		Docs:    "No changes will be applied",
		Allowed: true,
	},
	{
		Id:      "aikars",
		Name:    "Aikars",
		Docs:    "A popular set of tuned arguments for better performance. This works better for servers that have a lot of available memory (8GB+) and is not recommended otherwise",
		Allowed: true,
	},
	{
		Id:      "krusic",
		Name:    "Krusic",
		Docs:    "Another set of performance arguments",
		Allowed: true,
	},
	{
		Id:      "obydux",
		Name:    "Obydux",
		Docs:    "Another set of performance arguments",
		Allowed: true,
	},
}

var SnapshotsStorageOptions []Choice = []Choice{
	{
		Id:      "archive",
		Name:    "Archive",
		Docs:    "Store snapshots as an archive",
		Allowed: true,
	},
	{
		Id:      "folder",
		Name:    "Folder",
		Docs:    "Store snapshots as a folder",
		Allowed: true,
	},
}

var PackagePermissionOptions []Choice = []Choice{
	{
		Id:      "standard",
		Name:    "Standard",
		Docs:    "Default package permissions",
		Allowed: true,
	},
	{
		Id:      "restricted",
		Name:    "Restricted",
		Docs:    "Packages you do not trust",
		Allowed: true,
	},
	{
		Id:      "elevated",
		Name:    "Elevated",
		Docs:    "Packages that you trust and want to provide access to special commands for",
		Allowed: true,
	},
}

var CachingStrategyOptions []Choice = []Choice{
	{
		Id:      "all",
		Name:    "All",
		Docs:    "Cache all packages whenever you run the `package sync` command",
		Allowed: true,
	},
	{
		Id:      "none",
		Name:    "None",
		Docs:    "Never cache any scripts",
		Allowed: true,
	},
	{
		Id:      "lazy",
		Name:    "Lazy",
		Docs:    "Cache only when a package is requested",
		Allowed: true,
	},
}
