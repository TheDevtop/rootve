package main

import (
	"os/exec"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/TheDevtop/rootve/pkg/libve"
	"golang.org/x/sys/unix"
)

// A "Virtual Machine" structure
type vmach struct {
	config libve.VirtConfig
	state  byte
	proc   *exec.Cmd
}

// Allocate and initialize a "Virtual Machine"
func newVmach(name string, vc libve.VirtConfig) *vmach {
	newMach := new(vmach)
	newMach.config = vc
	newMach.state = libcsrv.StateOff
	newMach.proc = exec.Command(libcsrv.RootexecPath, libcsrv.RootexecFlagName, name)
	return newMach
}

// Executes the state switch function
func (vmp *vmach) Switch(state byte) error {
	var err error

	switch vmp.state {
	case libcsrv.StateOff:
		if state == libcsrv.StateOn {
			if err = vmp.proc.Start(); err != nil {
				return err
			}
			vmp.state = state
			return nil
		}

	case libcsrv.StateOn:
		switch state {
		case libcsrv.StateOff:
			if err = vmp.proc.Cancel(); err != nil {
				return err
			}
			vmp.state = state
			return nil
		case libcsrv.StatePaused:
			if err = vmp.proc.Process.Signal(unix.SIGTSTP); err != nil {
				return err
			}
			vmp.state = state
			return nil
		}

	case libcsrv.StatePaused:
		switch state {
		case libcsrv.StateOff:
			if err = vmp.proc.Cancel(); err != nil {
				return err
			}
			vmp.state = state
			return nil
		case libcsrv.StateOn:
			if err = vmp.proc.Process.Signal(unix.SIGCONT); err != nil {
				return err
			}
			vmp.state = state
			return nil
		}
	}
	return nil
}
