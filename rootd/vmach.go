package main

import (
	"errors"
	"os/exec"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/TheDevtop/rootve/pkg/libve"
	"golang.org/x/sys/unix"
)

var errVmapEntry = errors.New("vmap entry not found")

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
	newMach.proc.SysProcAttr = &unix.SysProcAttr{
		Setpgid: true,
	}
	return newMach
}

// Safely stop a "Virtual Machine"
func (vmp *vmach) safeStop() {
	unix.Kill(vmp.proc.Process.Pid, unix.SIGKILL)
	vmp.proc.Process.Release()
}

// Executes the state switch function
func (vmp *vmach) Switch(state byte) error {
	var (
		err      error
		stateErr = errors.New("invalid state transition")
	)

	switch vmp.state {
	case libcsrv.StateOff:
		if state == libcsrv.StateOn {
			if err = vmp.proc.Start(); err != nil {
				return err
			}
			vmp.state = state
			return nil
		}
		return stateErr
	case libcsrv.StateOn:
		switch state {
		case libcsrv.StateOff:
			vmp.safeStop()
			vmp.state = state
			return nil
		case libcsrv.StatePaused:
			if err = vmp.proc.Process.Signal(unix.SIGTSTP); err != nil {
				return err
			}
			vmp.state = state
			return nil
		}
		return stateErr
	case libcsrv.StatePaused:
		switch state {
		case libcsrv.StateOff:
			vmp.safeStop()
			vmp.state = state
			return nil
		case libcsrv.StateOn:
			if err = vmp.proc.Process.Signal(unix.SIGCONT); err != nil {
				return err
			}
			vmp.state = state
			return nil
		}
		return stateErr
	}
	return stateErr
}
