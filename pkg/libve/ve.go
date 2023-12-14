package libve

import (
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

type VirtEnv struct {
	root string
	proc exec.Cmd
	uid  int
	gid  int
}

// Change root and directory
func (ve *VirtEnv) Chroot() error {
	var err error
	if err = unix.Chroot(ve.root); err != nil {
		return err
	}
	if err = unix.Chdir("/"); err != nil {
		return err
	}
	return nil
}

// Execute the process inside
func (ve *VirtEnv) Execute() error {
	return ve.proc.Run()
}

// Attach the standard devices
func (ve *VirtEnv) Attach(in, out, err *os.File) {
	ve.proc.Stdin = in
	ve.proc.Stdout = out
	ve.proc.Stderr = err
}

// Set the user and group id
func (ve *VirtEnv) SetCreds() error {
	var err error
	if unix.Setuid(ve.uid); err != nil {
		return err
	}
	if unix.Setgid(ve.gid); err != nil {
		return err
	}
	return nil
}

// Attempt to mount all filesystems
func (ve *VirtEnv) Mount() {
	exec.Command("/sbin/mount", "-a").Run()
}

// Allocate virtual environment
func NewEnvironment(vc VirtConfig) *VirtEnv {
	ve := new(VirtEnv)
	ve.proc = *new(exec.Cmd)

	ve.root = vc.Root
	ve.proc.Dir = vc.Directory
	ve.uid = vc.Uid
	ve.gid = vc.Gid
	ve.proc.Env = vc.Environment
	ve.proc.Path = vc.CommandPath
	ve.proc.Args = vc.CommandArgs
	return ve
}
