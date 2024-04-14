package profile_schema

type Rcon struct {
	Enable   bool   `json:"enable,omitempty"`
	Port     int    `json:"port,omitempty"`
	Password string `json:"password,omitempty"`
}

type Query struct {
	Enable bool `json:"enable,omitempty"`
	Port   int  `json:"port,omitempty"`
}

type Whitelist struct {
	Enable  bool `json:"enable,omitempty"`
	Enforce bool `json:"enforce,omitempty"`
}

type Gamemode struct {
	Default string `json:"default,omitempty"`
	Force   bool   `json:"force,omitempty"`
}

type Datapacks struct {
	FunctionPermissionLevel int      `json:"function_permission_level,omitempty"`
	InitialEnabled          []string `json:"initial_enabled,omitempty"`
	InitialDisabled         []string `json:"initial_disabled,omitempty"`
}

type World struct {
	Name              string      `json:"name,omitempty"`
	Seed              string      `json:"seed,omitempty"`
	Type              string      `json:"type,omitempty"`
	Structures        bool        `json:"structures,omitempty"`
	GeneratorSettings interface{} `json:"generator_settings,omitempty"`
	MaxSize           int         `json:"max_size,omitempty"`
	MaxBuildHeight    int         `json:"max_build_height,omitempty"`
	AllowNether       bool        `json:"allow_nether,omitempty"`
}

type ResourcePack struct {
	Uri      string `json:"uri,omitempty"`
	Prompt   string `json:"prompt,omitempty"`
	Sha1     string `json:"sha1,omitempty"`
	Required bool   `json:"required,omitempty"`
}

type ServerOptions struct {
	Rcon                        Rcon         `json:"rcon,omitempty"`
	Query                       Query        `json:"query,omitempty"`
	Whitelist                   Whitelist    `json:"whitelist,omitempty"`
	Gamemode                    Gamemode     `json:"gamemode,omitempty"`
	Datapacks                   Datapacks    `json:"datapacks,omitempty"`
	World                       World        `json:"world,omitempty"`
	ResourcePack                ResourcePack `json:"resource_pack,omitempty"`
	Custom                      interface{}  `json:"custom,omitempty"`
	AllowFlight                 bool         `json:"allow_flight,omitempty"`
	BroadcastConsoleToOps       bool         `json:"broadcast_console_to_ops,omitempty"`
	BroadcastRconToOps          bool         `json:"broadcast_rcon_to_ops,omitempty"`
	Difficulty                  interface{}  `json:"difficulty,omitempty"`
	AllowCommandBlocks          bool         `json:"allow_command_blocks,omitempty"`
	JmxMonitoring               bool         `json:"jmx_monitoring,omitempty"`
	EnableStatus                bool         `json:"enable_status,omitempty"`
	EnforceSecureProfile        bool         `json:"enforce_secure_profile,omitempty"`
	EntityBroadcastRange        int          `json:"entity_broadcast_range,omitempty"`
	MaxChainedNeighborUpdates   int          `json:"max_chained_neighbor_updates,omitempty"`
	MaxPlayers                  int          `json:"max_players,omitempty"`
	MaxTickTime                 int          `json:"max_tick_time,omitempty"`
	Motd                        string       `json:"motd,omitempty"`
	NetworkCompressionThreshold interface{}  `json:"network_compression_threshold,omitempty"`
	OfflineMode                 bool         `json:"offline_mode,omitempty"`
	OpPermissionLevel           int          `json:"op_permission_level,omitempty"`
	PlayerIdleTimeout           int          `json:"player_idle_timeout,omitempty"`
	PreventProxyConnections     bool         `json:"prevent_proxy_connections,omitempty"`
	EnableChatPreview           bool         `json:"enable_chat_preview,omitempty"`
	EnablePvp                   bool         `json:"enable_pvp,omitempty"`
	RateLimit                   int          `json:"rate_limit,omitempty"`
	Ip                          string       `json:"ip,omitempty"`
	Port                        int          `json:"port,omitempty"`
	SimulationDistance          int          `json:"simulation_distance,omitempty"`
	EnableSnooper               bool         `json:"enable_snooper,omitempty"`
	SpawnAnimals                bool         `json:"spawn_animals,omitempty"`
	SpawnMonsters               bool         `json:"spawn_monsters,omitempty"`
	SpawnNpcs                   bool         `json:"spawn_npcs,omitempty"`
	SpawnProtection             int          `json:"spawn_protection,omitempty"`
	SyncChunkWrites             bool         `json:"sync_chunk_writes,omitempty"`
	UseNativeTransport          bool         `json:"use_native_transport,omitempty"`
	ViewDistance                int          `json:"view_distance,omitempty"`
}
