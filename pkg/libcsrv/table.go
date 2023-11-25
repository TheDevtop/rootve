package libcsrv

import (
	"os/exec"

	"github.com/TheDevtop/rootve/pkg/libve"
)

type VeTable map[string]*VeEntry

// Make virtual environment table out of configuration structure
func MakeTable(mvc map[string]libve.VirtConfig) VeTable {
	vtab := make(VeTable, len(mvc))

	for key, val := range mvc {
		entry := new(VeEntry)

		entry.State = StateOff
		entry.Config = val
		entry.Exec = *exec.Command(RootexecPath, RootexecFlagName, key)

		vtab[key] = entry
	}
	return vtab
}
