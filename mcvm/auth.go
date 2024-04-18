package mcvm

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bbfh-dev/configure.mcvm/cli"
	"github.com/bbfh-dev/configure.mcvm/mcvm/schema/auth"
)

type Auth struct {
	Users map[string]auth.AuthUser `json:"users,omitempty"`
}

func (config *Auth) Decode(data []byte) error {
	err := json.Unmarshal(data, &config)
	if err != nil {
		cli.Logger(fmt.Sprintf("JSON Decode: %s", err)).Error()
		return err
	}

	return nil
}

func (config *Auth) DecodeFromFile() error {
	file, err := os.ReadFile(GetDataFile("internal", "auth", "db.json"))
	if err != nil {
		return err
	}

	return config.Decode(file)
}

func (config *Auth) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(&config, "", "  ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (config *Auth) EncodeToFile() error {
	data, err := config.Encode()
	if err != nil {
		return nil
	}

	return os.WriteFile(GetDataFile("internal", "auth", "db.json"), data, 0644)
}
