package libve

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

const (
	DefaultPath = "/etc/rootve"
)

type VirtConfig struct {
	Root        string
	Directory   string
	Clean       bool
	CommandPath string
	CommandArgs []string
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
	if err = toml.Unmarshal(buf, mvc); err != nil {
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
