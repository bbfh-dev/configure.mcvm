package mcvm

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bbfh-dev/configure.mcvm/cli"
	"github.com/bbfh-dev/configure.mcvm/mcvm/schema"
	profile_schema "github.com/bbfh-dev/configure.mcvm/mcvm/schema/profile"
)

type Config struct {
	Users           map[string]schema.User        `json:"users,omitempty"`
	DefaultUser     string                        `json:"default_user,omitempty"`
	Profiles        map[string]schema.Profile     `json:"profiles,omitempty"`
	Packages        []profile_schema.JsonPackages `json:"packages,omitempty"`
	InstancePresets interface{}                   `json:"instance_presets,omitempty"` // Ignore
	Preferences     schema.Preferences            `json:"preferences,omitempty"`
}

func (config *Config) Decode(data []byte) error {
	err := json.Unmarshal(data, &config)
	if err != nil {
		cli.Logger(fmt.Sprintf("JSON Decode: %s", err)).Error()
		return err
	}

	return nil
}

func (config *Config) DecodeFromFile() error {
	file, err := os.ReadFile(GetConfigFile("mcvm.json"))
	if err != nil {
		return err
	}

	return config.Decode(file)
}

func (config *Config) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(&config, "", "  ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (config *Config) EncodeToFile() error {
	data, err := config.Encode()
	if err != nil {
		return nil
	}

	return os.WriteFile(GetConfigFile("mcvm.json"), data, 0644)
}
