package schema

import profile_schema "github.com/bbfh-dev/configure.mcvm/mcvm/schema/profile"

type Profile struct {
	Version          string                                 `json:"version,omitempty"`
	ModLoader        *string                                `json:"modloader,omitempty"`
	ClientType       *string                                `json:"client_type,omitempty"`
	ServerType       *string                                `json:"server_type,omitempty"`
	Instances        map[string]profile_schema.JsonInstance `json:"instances,omitempty"`
	Packages         *profile_schema.JsonPackages           `json:"packages,omitempty"`
	PackageStability *string                                `json:"package_stability,omitempty"`
}
