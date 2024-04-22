package config

import (
	"encoding/json"
	"os"

	"github.com/bbfh-dev/configure.mcvm/cli"
)

var MCVMConfig Config

type Config struct {
	Users       map[string]User `json:"users,omitempty"`
	DefaultUser string          `json:"default_user,omitempty"`
}

func (cfg *Config) Decode(data []byte) error {
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		cli.Log("JSON Decode ERROR: %s", err)
		return err
	}

	return nil
}

func (cfg *Config) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(&cfg, "", "\t")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (cfg *Config) DecodeFromFile(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return cfg.Decode(file)
}

func (cfg *Config) EncodeToFile(path string) error {
	data, err := cfg.Encode()
	if err != nil {
		return nil
	}

	return os.WriteFile(path, data, 0644)
}
