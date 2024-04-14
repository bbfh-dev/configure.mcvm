package mcvm_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/bbfh-dev/configure.mcvm/mcvm"
	"github.com/bbfh-dev/configure.mcvm/mcvm/schema"
	profile_schema "github.com/bbfh-dev/configure.mcvm/mcvm/schema/profile"
)

func readFile(t *testing.T, name string) []byte {
	data, err := os.ReadFile(filepath.Join("../", "data", name))
	if err != nil {
		t.Fatal(err)
	}

	return data
}

var modloader = "fabric"
var dir = "/tmp/"
var stability = "stable"
var defaultMCVM = mcvm.Config{Profiles: map[string]schema.Profile{
	"test": {
		Version:    "1.20.4",
		ModLoader:  &modloader,
		ClientType: &modloader,
		ServerType: &modloader,
		Instances: map[string]profile_schema.JsonInstance{
			"example-server": {Instance: profile_schema.Instance{Type: "server"}},
			"example-client": {Instance: profile_schema.Instance{Type: "client"}},
			"example": {
				Instance: profile_schema.Instance{Type: "client", DatapackFolder: &dir},
			},
		},
		Packages: &profile_schema.JsonPackages{PackageSet: profile_schema.PackageSet{
			Global: []profile_schema.Package{
				{
					ID:   "fabric-api",
					Type: "repository",
				},
			},
			Client: []profile_schema.Package{},
			Server: []profile_schema.Package{},
		}},
		PackageStability: &stability,
	},
}}

func TestDecode(t *testing.T) {
	expected := defaultMCVM
	expected_json, _ := json.MarshalIndent(expected, "", "  ")

	// Full
	got := mcvm.Config{}
	got.Decode(readFile(t, "full_mcvm.json"))
	got_json, _ := json.MarshalIndent(got, "", "  ")

	if reflect.DeepEqual(got, expected) {
		t.Fatalf("FULL # GOT:\n %s \n# EXPECTED:\n %s", got_json, expected_json)
	}

	// Mixed
	got = mcvm.Config{}
	err := got.Decode(readFile(t, "mixed_mcvm.json"))
	if err != nil {
		t.Fatal(err)
	}
	got_json, _ = json.MarshalIndent(got, "", "  ")

	if reflect.DeepEqual(got, expected) {
		t.Fatalf("MIXED # GOT:\n %s \n# EXPECTED:\n %s", got_json, expected_json)
	}

	// Short
	got = mcvm.Config{}
	err = got.Decode(readFile(t, "short_mcvm.json"))
	if err != nil {
		t.Fatal(err)
	}
	got_json, _ = json.MarshalIndent(got, "", "  ")

	if reflect.DeepEqual(got, expected) {
		t.Fatalf("SHORT # GOT:\n %s \n# EXPECTED:\n %s", got_json, expected_json)
	}
}

// DEBUGGING:
// func TestEncode(t *testing.T) {
// 	str := mcvm.Config{}
// 	err := str.Decode(readFile(t, "full_mcvm.json"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	data, err := str.Encode()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	t.Fatalf("%s", data)
// }
