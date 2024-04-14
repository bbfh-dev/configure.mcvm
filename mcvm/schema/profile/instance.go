package profile_schema

import (
	"encoding/json"
	"fmt"

	"github.com/bbfh-dev/configure.mcvm/mcvm/schema/marshal"
)

type JsonInstance struct {
	Instance
}

func (instance *JsonInstance) UnmarshalJSON(data []byte) error {
	if marshal.IsString(data[0]) {
		var newInstance string
		if err := json.Unmarshal(data, &newInstance); err != nil {
			return err
		}

		instance.Instance = Instance{Type: newInstance}

		return nil
	}

	return json.Unmarshal(data, &instance.Instance)
}

func (instance JsonInstance) MarshalJSON() ([]byte, error) {
	fields := marshal.GetFilledFields(instance.Instance)
	if len(fields) == 1 && fields[0] == "Type" {
		return []byte(fmt.Sprintf("%q", instance.Instance.Type)), nil
	}

	return json.Marshal(instance.Instance)
}

type LaunchArgs struct {
	Jvm  []string `json:"jvm,omitempty"`
	Game []string `json:"game,omitempty"`
}

type Memory struct {
	Init string `json:"init,omitempty"`
	Max  string `json:"max,omitempty"`
}

type QuickPlay struct {
	Type   string `json:"type,omitempty"` // "world" | "server" | "realm"
	World  string `json:"world,omitempty"`
	Server string `json:"server,omitempty"`
	Port   string `json:"port,omitempty"`
	Realm  string `json:"realm,omitempty"`
}

type Wrapper struct {
	Cmd  string   `json:"cmd,omitempty"`
	Args []string `json:"args,omitempty"`
}

type Launch struct {
	Args           LaunchArgs        `json:"args,omitempty"`
	Memory         interface{}       `json:"memory,omitempty"` // TODO: Could be either string or Memory
	Env            map[string]string `json:"env,omitempty"`
	Wrapper        Wrapper           `json:"wrapper,omitempty"`
	Java           string            `json:"java,omitempty"`   // "auto" | "system" | "adoptium" | "zulu" | "graalvm" | string
	Preset         string            `json:"preset,omitempty"` // "none" | "akairs" | "krusic" | "obydux"
	QuickPlay      QuickPlay         `json:"quick_play,omitempty"`
	UseLog4jConfig bool              `json:"use_log4j_config,omitempty"`
}

type WindowResolution struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Window struct {
	Resolution WindowResolution `json:"resolution,omitempty"`
}

type Snapshots struct {
	Paths       []string `json:"paths,omitempty"`
	MaxCount    int      `json:"max_count,omitempty"`
	StorageType string   `json:"storage_type,omitempty"`
}

type TypeOptions struct {
	Type string `json:"type,omitempty"`
}

type Instance struct {
	Type           string       `json:"type"`
	Launch         *Launch      `json:"launch,omitempty"`
	Options        *JsonOptions `json:"options,omitempty"`
	Window         *Window      `json:"window,omitempty"`
	DatapackFolder *string      `json:"datapack_folder,omitempty"`
	Snapshots      *Snapshots   `json:"snapshots,omitempty"`
	Packages       *[]Package   `json:"packages,omitempty"`
	Preset         *string      `json:"preset,omitempty"`
}

type JsonOptions struct {
	Options interface{}
}

func (options *JsonOptions) UnmarshalJSON(data []byte) error {
	var clientOptions ClientOptions
	if err := json.Unmarshal(data, &clientOptions); err == nil {
		if clientOptions.DataVersion != 0 {
			options.Options = clientOptions
			return nil
		}
	}

	var serverOptions ServerOptions
	err := json.Unmarshal(data, &serverOptions)
	options.Options = serverOptions

	return err
}
