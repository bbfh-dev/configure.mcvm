package profile_schema

type Video struct {
	Vsync                  bool                 `json:"vsync,omitempty"`
	EntityShadows          bool                 `json:"entity_shadows,omitempty"`
	Fullscreen             bool                 `json:"fullscreen,omitempty"`
	ViewBobbing            bool                 `json:"view_bobbing,omitempty"`
	DarkMojangBackground   bool                 `json:"dark_mojang_background,omitempty"`
	HideLightningFlashes   bool                 `json:"hide_lightning_flashes,omitempty"`
	Fov                    int                  `json:"fov,omitempty"`
	ScreenEffectScale      float64              `json:"screen_effect_scale,omitempty"`
	FovEffectScale         float64              `json:"fov_effect_scale,omitempty"`
	DarknessEffectScale    float64              `json:"darkness_effect_scale,omitempty"`
	Brightness             float64              `json:"brightness,omitempty"`
	RenderDistance         int                  `json:"render_distance,omitempty"`
	SimulationDistance     int                  `json:"simulation_distance,omitempty"`
	EntityDistanceScaling  float64              `json:"entity_distance_scaling,omitempty"`
	GuiScale               float64              `json:"gui_scale,omitempty"`
	Particles              interface{}          `json:"particles,omitempty"`
	MaxFPS                 float64              `json:"max_fps,omitempty"`
	GraphicsMode           interface{}          `json:"graphics_mode,omitempty"`
	SmoothLighting         bool                 `json:"smooth_lighting,omitempty"`
	ChunkUpdatesMode       interface{}          `json:"chunk_updates_mode,omitempty"`
	BiomeBlend             int                  `json:"biome_blend,omitempty"`
	Clouds                 interface{}          `json:"clouds,omitempty"`
	MipmapLevels           int                  `json:"mipmap_levels,omitempty"`
	WindowWidth            int                  `json:"window_width,omitempty"`
	WindowHeight           int                  `json:"window_height,omitempty"`
	AttackIndicator        interface{}          `json:"attack_indicator,omitempty"`
	FullscreenResolution   FullscreenResolution `json:"fullscreen_resolution,omitempty"`
	AllowBlockAlternatives bool                 `json:"allow_block_alternatives,omitempty"`
}

type FullscreenResolution struct {
	Width       int `json:"width,omitempty"`
	Height      int `json:"height,omitempty"`
	RefreshRate int `json:"refresh_rate,omitempty"`
	ColorBits   int `json:"color_bits,omitempty"`
}

type Keybind struct {
	Key     string `json:"key,omitempty"`
	Alt     bool   `json:"alt,omitempty"`
	Ctrl    bool   `json:"ctrl,omitempty"`
	Shift   bool   `json:"shift,omitempty"`
	KeyCode int    `json:"key_code,omitempty"`
}

type Control struct {
	Keys                  map[string]Keybind `json:"keys,omitempty"`
	AutoJump              bool               `json:"auto_jump,omitempty"`
	InvertMouseY          bool               `json:"invert_mouse_y,omitempty"`
	EnableTouchscreen     bool               `json:"enable_touchscreen,omitempty"`
	ToggleSprint          bool               `json:"toggle_sprint,omitempty"`
	ToggleCrouch          bool               `json:"toggle_crouch,omitempty"`
	MouseSensitivity      int                `json:"mouse_sensitivity,omitempty"`
	MouseWheelSensitivity float64            `json:"mouse_wheel_sensitivity,omitempty"`
	RawMouseInput         bool               `json:"raw_mouse_input,omitempty"`
}

type Chat struct {
	AutoCommandSuggestions bool        `json:"auto_command_suggestions,omitempty"`
	EnableColors           bool        `json:"enable_colors,omitempty"`
	EnableLinks            bool        `json:"enable_links,omitempty"`
	PromptLinks            bool        `json:"prompt_links,omitempty"`
	ForceUnicode           bool        `json:"force_unicode,omitempty"`
	Visibility             interface{} `json:"visibility,omitempty"`
	Opacity                float64     `json:"opacity,omitempty"`
	LineSpacing            float64     `json:"line_spacing,omitempty"`
	BackgroundOpacity      float64     `json:"background_opacity,omitempty"`
	BackgroundForChatOnly  bool        `json:"background_for_chat_only,omitempty"`
	FocusedHeight          float64     `json:"focused_height,omitempty"`
	UnfocusedHeight        float64     `json:"unfocused_height,omitempty"`
	Delay                  float64     `json:"delay,omitempty"`
	Scale                  float64     `json:"scale,omitempty"`
	Width                  float64     `json:"width,omitempty"`
	NarratorMode           interface{} `json:"narrator_mode,omitempty"`
}

type Volume struct {
	Master  float64 `json:"master,omitempty"`
	Music   float64 `json:"music,omitempty"`
	Record  float64 `json:"record,omitempty"`
	Weather float64 `json:"weather,omitempty"`
	Block   float64 `json:"block,omitempty"`
	Hostile float64 `json:"hostile,omitempty"`
	Neutral float64 `json:"neutral,omitempty"`
	Player  float64 `json:"player,omitempty"`
	Ambient float64 `json:"ambient,omitempty"`
	Voice   float64 `json:"voice,omitempty"`
}

type Sound struct {
	Volume           Volume `json:"volume,omitempty"`
	ShowSubtitles    bool   `json:"show_subtitles,omitempty"`
	DirectionalAudio bool   `json:"directional_audio,omitempty"`
	Device           string `json:"device,omitempty"`
}

type Skin struct {
	Cape        bool `json:"cape,omitempty"`
	Jacket      bool `json:"jacket,omitempty"`
	LeftSleeve  bool `json:"left_sleeve,omitempty"`
	RightSleeve bool `json:"right_sleeve,omitempty"`
	LeftPants   bool `json:"left_pants,omitempty"`
	RightPants  bool `json:"right_pants,omitempty"`
	Hat         bool `json:"hat,omitempty"`
}

type Stream struct {
	BytesPerPixel            float64 `json:"bytes_per_pixel,omitempty"`
	ChatEnabled              bool    `json:"chat_enabled,omitempty"`
	ChatFilter               bool    `json:"chat_filter,omitempty"`
	Compression              bool    `json:"compression,omitempty"`
	FPS                      float64 `json:"fps,omitempty"`
	Bitrate                  float64 `json:"bitrate,omitempty"`
	MicrophoneToggleBehavior bool    `json:"microphone_toggle_behavior,omitempty"`
	MicrophoneVolume         float64 `json:"microphone_volume,omitempty"`
	PreferredServer          string  `json:"preferred_server,omitempty"`
	SendMetadata             bool    `json:"send_metadata,omitempty"`
	SystemVolume             float64 `json:"system_volume,omitempty"`
}

type ClientOptions struct {
	DataVersion            int         `json:"data_version,omitempty"`
	Video                  Video       `json:"video,omitempty"`
	Control                Control     `json:"control,omitempty"`
	Chat                   Chat        `json:"chat,omitempty"`
	Sound                  Sound       `json:"sound,omitempty"`
	Skin                   Skin        `json:"skin,omitempty"`
	Stream                 Stream      `json:"stream,omitempty"`
	Custom                 interface{} `json:"custom,omitempty"`
	RealmsNotifications    bool        `json:"realms_notifications,omitempty"`
	ReducedDebugInfo       bool        `json:"reduced_debug_info,omitempty"`
	Difficulty             interface{} `json:"difficulty,omitempty"`
	ResourcePacks          []string    `json:"resource_packs,omitempty"`
	Language               string      `json:"language,omitempty"`
	TutorialStep           interface{} `json:"tutorial_step,omitempty"`
	SkipMultiplayerWarning bool        `json:"skip_multiplayer_warning,omitempty"`
	SkipRealms32BitWarning bool        `json:"skip_realms_32_bit_warning,omitempty"`
	HideBundleTutorial     bool        `json:"hide_bundle_tutorial,omitempty"`
	JoinedServer           bool        `json:"joined_server,omitempty"`
	SyncChunkWrites        bool        `json:"sync_chunk_writes,omitempty"`
	UseNativeTransport     bool        `json:"use_native_transport,omitempty"`
	HeldItemTooltips       bool        `json:"held_item_tooltips,omitempty"`
	AdvancedItemTooltips   bool        `json:"advanced_item_tooltips,omitempty"`
	LogLevel               interface{} `json:"log_level,omitempty"`
	HideMatchedNames       bool        `json:"hide_matched_names,omitempty"`
	PauseOnLostFocus       bool        `json:"pause_on_lost_focus,omitempty"`
	MainHand               interface{} `json:"main_hand,omitempty"`
	HideServerAddress      bool        `json:"hide_server_address,omitempty"`
	ShowAutosaveIndicator  bool        `json:"show_autosave_indicator,omitempty"`
	AllowServerListing     bool        `json:"allow_server_listing,omitempty"`
}
