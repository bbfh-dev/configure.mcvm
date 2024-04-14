package profile_schema

import (
	"encoding/json"

	"github.com/bbfh-dev/configure.mcvm/mcvm/schema/marshal"
)

type JsonPackages struct {
	PackageSet
}

func (packages *JsonPackages) UnmarshalJSON(data []byte) error {
	if marshal.IsArray(data[0]) {
		decoder := marshal.NewDecoder(&data)
		decoder.PrepareArray()
		result := make([]Package, decoder.PrepareArray())

		var tryString string
		var tryPackage Package
		err := decoder.Iterate(
			marshal.ItemDecoder{
				Value: &tryString,
				Callback: func(i int) {
					result[i] = Package{ID: tryString}
				},
			},
			marshal.ItemDecoder{
				Value: &tryPackage,
				Callback: func(i int) {
					result[i] = tryPackage
				},
			},
		).Error()

		packages.PackageSet = PackageSet{
			Global: result,
		}

		return err
	}

	return json.Unmarshal(data, &packages.PackageSet)
}

type Package struct {
	ID                 string    `json:"id,omitempty"`
	Type               string    `json:"type,omitempty"`
	Features           *[]string `json:"features,omitempty"`
	UseDefaultFeatures *bool     `json:"use_default_features,omitempty"`
	Permissions        *string   `json:"permissions,omitempty"`
	Stability          *string   `json:"stability,omitempty"`
	Worlds             *[]string `json:"worlds,omitempty"`
}

type PackageSet struct {
	Global []Package `json:"global,omitempty"`
	Client []Package `json:"client,omitempty"`
	Server []Package `json:"server,omitempty"`
}
