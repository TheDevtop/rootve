package libve

import (
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

// The "Virtual Environment"
type VirtEnv struct {
	root   string
	proc   exec.Cmd
	uid    int
	gid    int
	net    bool
	netbr  string
	netif  string
	addrv4 string
	maskv4 string
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

// Set the process to be a shell
func (ve *VirtEnv) SetShell() {
	ve.proc.Path = "/bin/ksh"
	ve.proc.Args = []string{"-l"}
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
func (ve *VirtEnv) Mount() error {
	return exec.Command("/sbin/mount", "-a").Run()
}

// Attempt to unmount all filesystems
func (ve *VirtEnv) Umount() error {
	return exec.Command("/sbin/umount", "-a").Run()
}

// Configure the standard devices
func (ve *VirtEnv) Stdinit(shellMode bool) error {
	ve.proc.Stdin = os.Stdin
	ve.proc.Stdout = os.Stdout
	ve.proc.Stderr = os.Stderr
	if !shellMode {
		return unix.Setpgid(os.Getpid(), 0)
	}
	return nil
}

// Attempt to initialize devices
func (ve *VirtEnv) Devinit() {
	devcmd := exec.Command("/dev/MAKEDEV", "std", "fd", "ptm", "tty0", ve.netif)
	devcmd.Dir = "/dev/"
	devcmd.Run()
}

// Initialize networking
func (ve *VirtEnv) Netinit() error {
	var (
		err error
		cmd *exec.Cmd
	)

	if !ve.net {
		return nil
	}

	// Create network interface
	cmd = exec.Command("/sbin/ifconfig", ve.netif, "create")
	cmd.Run()

	// Add interface to bridge
	cmd = exec.Command("/sbin/brconfig", ve.netbr, "add", ve.netif)
	cmd.Run()

	// Configure interface with address and netmask
	cmd = exec.Command("/sbin/ifconfig", ve.netif, "inet", ve.addrv4, "netmask", ve.maskv4)
	if err = cmd.Run(); err != nil {
		return err
	}

	// Turn the interface on
	cmd = exec.Command("/sbin/ifconfig", ve.netif, "up")
	if err = cmd.Run(); err != nil {
		return err
	}

	return nil
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
	ve.net = vc.Networking
	ve.netbr = vc.Bridge
	ve.netif = vc.Interface
	ve.addrv4 = vc.AddressV4
	ve.maskv4 = vc.NetmaskV4

	return ve
}
