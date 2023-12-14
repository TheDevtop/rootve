package librex

import (
	"os/exec"

	"github.com/TheDevtop/rootve/pkg/libve"
	"golang.org/x/sys/unix"
)

// Rootexec instance structure
type Rex struct {
	Config libve.VirtConfig
	State  byte
	proc   *exec.Cmd
}

// Start and execute rootexec instance
func (rexPtr *Rex) Start() error {
	if err := rexPtr.proc.Start(); err != nil {
		return err
	}
	rexPtr.State = StateOn
	return nil
}

// Stop rootexec instance
func (rexPtr *Rex) Stop() error {
	if err := rexPtr.proc.Process.Kill(); err != nil {
		return err
	}
	if err := rexPtr.proc.Process.Release(); err != nil {
		return err
	}
	rexPtr.State = StateOff
	return nil
}

// Pause rootexec instance

// Resume rootexec instance

// Allocate rootexec instance structure
func NewRex(name string, vc libve.VirtConfig) *Rex {
	rexPtr := new(Rex)
	rexPtr.Config = vc
	rexPtr.State = StateOff
	rexPtr.proc = exec.Command(RootexecPath, RootexecFlagName, name)
	rexPtr.proc.SysProcAttr = &unix.SysProcAttr{
		Setpgid: true,
	}
	return rexPtr
}
