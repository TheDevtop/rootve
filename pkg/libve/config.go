package libve

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Default configuration file path
const ConfigPath = "/etc/rootve"

// Configuration entry structure
type VirtConfig struct {
	// Environmental
	Root        string
	Autoboot    bool
	Directory   string
	Uid         int
	Gid         int
	Environment []string

	// Client process
	CommandPath string
	CommandArgs []string

	// Networking
	Networking bool
	Bridge     string
	Interface  string
	AddressV4  string
}

// Read configuration map from toml file
func ReadConfig(path string) (map[string]VirtConfig, error) {
	var (
		mvc = make(map[string]VirtConfig)
		err error
		buf []byte
	)

	if buf, err = os.ReadFile(path); err != nil {
		return nil, err
	}
	if err = toml.Unmarshal(buf, &mvc); err != nil {
		return nil, err
	}
	return mvc, nil
}

// Write configuration map to toml file
func WriteConfig(path string, mvc map[string]VirtConfig) error {
	var (
		err error
		buf []byte
	)

	if buf, err = toml.Marshal(mvc); err != nil {
		return err
	}
	os.WriteFile(path, buf, 0660)
	return nil
}

// Allocate and initialize VE configuration
// with default values
func MakeVirtConfig() VirtConfig {
	return VirtConfig{
		Root:        "/",
		Autoboot:    false,
		Directory:   "/",
		Uid:         0,
		Gid:         0,
		Environment: []string{"TERM=xterm"},
		CommandPath: "/bin/ksh",
		CommandArgs: []string{"-l"},
		Networking:  false,
		Bridge:      "br0",
		Interface:   "lo",
		AddressV4:   "0.0.0.0",
	}
}
