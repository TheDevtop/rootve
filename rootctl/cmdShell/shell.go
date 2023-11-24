package cmdShell

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

const TagShell = "shell"

func ShellMain() {
	var (
		err error
		cmd *exec.Cmd
	)

	if len(os.Args) < 2 {
		fmt.Println("Usage: shell [name]")
		os.Exit(2)
	}

	cmd = exec.Command(libcsrv.RootexecPath, libcsrv.RootexecFlagName, os.Args[1], libcsrv.RootexecFlagOverride, "/bin/ksh", "-l")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	os.Exit(0)
}
