package libve

import (
	"os"
	"os/exec"
	"syscall"
)

type VirtEnv struct {
	root string
	proc exec.Cmd
}

func (ve *VirtEnv) Chroot() error {
	var err error
	if err = syscall.Chroot(ve.root); err != nil {
		return err
	}
	if err = syscall.Chdir("/"); err != nil {
		return err
	}
	return nil
}

func (ve *VirtEnv) Execute() error {
	return ve.proc.Run()
}

func (ve *VirtEnv) Attach(in, out, err *os.File) {
	ve.proc.Stdin = in
	ve.proc.Stdout = out
	ve.proc.Stderr = err
}

func NewEnvironment(vc VirtConfig) *VirtEnv {
	ve := new(VirtEnv)
	ve.root = vc.Root
	ve.proc = *new(exec.Cmd)
	if vc.Clean {
		ve.proc.Env = nil
	}
	ve.proc.Path = vc.CommandPath
	ve.proc.Args = vc.CommandArgs
	ve.proc.Dir = vc.Directory
	return ve
}
