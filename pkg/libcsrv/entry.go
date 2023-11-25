package libcsrv

import (
	"os/exec"

	"github.com/TheDevtop/rootve/pkg/libve"
	"golang.org/x/sys/unix"
)

type VeEntry struct {
	State  string
	Config libve.VirtConfig
	Exec   exec.Cmd
}

func (entry *VeEntry) Start() error {
	switch entry.State {
	case StateOn:
		return nil
	case StateOff:
		if err := entry.Exec.Start(); err != nil {
			return err
		}
		entry.State = StateOn
	case StatePaused:
		return errInvalidState
	}
	return nil
}

func (entry *VeEntry) Stop() error {
	switch entry.State {
	case StateOn:
		if err := entry.Exec.Cancel(); err != nil {
			return err
		}
		entry.State = StateOff
	case StateOff:
		return nil
	case StatePaused:
		return errInvalidState
	}
	return nil
}

func (entry *VeEntry) Pause() error {
	switch entry.State {
	case StateOn:
		if err := entry.Exec.Process.Signal(unix.SIGTSTP); err != nil {
			return err
		}
		entry.State = StatePaused
	case StateOff:
		return errInvalidState
	case StatePaused:
		return nil
	}
	return nil
}

func (entry *VeEntry) Resume() error {
	switch entry.State {
	case StateOn:
		return nil
	case StateOff:
		return errInvalidState
	case StatePaused:
		if err := entry.Exec.Process.Signal(unix.SIGCONT); err != nil {
			return err
		}
		entry.State = StateOn
	}
	return nil
}
