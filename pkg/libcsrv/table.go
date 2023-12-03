package libcsrv

import (
	"encoding/json"
	"io"
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

func ReadTable(rc io.ReadCloser) (VeTable, error) {
	var (
		err error
		buf []byte
		tab = make(VeTable)
	)

	defer rc.Close()

	if buf, err = io.ReadAll(rc); err != nil {
		return nil, err
	}

	if json.Unmarshal(buf, &tab); err != nil {
		return nil, err
	}

	return tab, nil
}
