package libcsrv

import (
	"os/exec"

	"github.com/TheDevtop/rootve/pkg/libve"
)

type VeEntry struct {
	State  string
	Config libve.VirtConfig
	Exec   exec.Cmd
}

type VeTable map[string]*VeEntry

// Make virtual environment table out of configuration structure
func MakeTable(mvc map[string]libve.VirtConfig) VeTable {
	vtab := make(VeTable, len(mvc))

	for key, val := range mvc {
		entry := new(VeEntry)

		entry.State = StateOff
		entry.Config = val
		entry.Exec = *exec.Command("/usr/local/bin/rootexec", "-n", key)

		vtab[key] = entry
	}
	return vtab
}
