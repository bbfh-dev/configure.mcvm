package auth

import (
	"encoding/json"
	"os"

	"github.com/bbfh-dev/configure.mcvm/cli"
)

// File at: ~/.local/share/mcvm/internal/auth/db.json
var MCVMAuth Auth

type Auth struct {
	Users map[string]AuthUser `json:"users,omitempty"`
}

type AuthUser struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}

func (cfg *Auth) Decode(data []byte) error {
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		cli.Log("JSON Decode ERROR: %s", err)
		return err
	}

	return nil
}

func (cfg *Auth) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(&cfg, "", "\t")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (cfg *Auth) DecodeFromFile(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return cfg.Decode(file)
}

func (cfg *Auth) EncodeToFile(path string) error {
	data, err := cfg.Encode()
	if err != nil {
		return nil
	}

	return os.WriteFile(path, data, 0644)
}
