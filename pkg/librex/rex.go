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

// Allocate rootexec instance structure
func NewRex(name string, vc libve.VirtConfig) *Rex {
	r := new(Rex)
	r.Config = vc
	r.State = StateOff
	r.proc = exec.Command(RootexecPath, RootexecFlagName, name)
	r.proc.SysProcAttr = &unix.SysProcAttr{
		Setpgid: true,
	}
	return r
}
