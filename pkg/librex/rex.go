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
	if err := unix.Kill(-rexPtr.proc.Process.Pid, unix.SIGKILL); err != nil {
		return err
	}
	rexPtr.proc.Process.Wait()
	rexPtr.State = StateOff
	return nil
}

// Pause rootexec instance
func (rexPtr *Rex) Pause() error {
	if err := unix.Kill(-rexPtr.proc.Process.Pid, unix.SIGSTOP); err != nil {
		return err
	}
	rexPtr.State = StatePaused
	return nil
}

// Resume rootexec instance
func (rexPtr *Rex) Resume() error {
	if err := unix.Kill(-rexPtr.proc.Process.Pid, unix.SIGCONT); err != nil {
		return err
	}
	rexPtr.State = StateOn
	return nil
}

// Allocate rootexec instance structure
func NewRex(name string, vc libve.VirtConfig) *Rex {
	rexPtr := new(Rex)
	rexPtr.Config = vc
	rexPtr.State = StateOff
	rexPtr.proc = exec.Command(RootexecPath, RootexecFlagName, name)
	return rexPtr
}
