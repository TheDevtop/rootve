package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/TheDevtop/rootve/pkg/librex"
)

const cmdShell = "shell"

func shellMain() int {
	var (
		err error
		cmd *exec.Cmd
	)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [name]\n", cmdShell)
		return 2
	}

	cmd = exec.Command(librex.RootexecPath, librex.RootexecFlagName, os.Args[1], librex.RootexecFlagAttach, librex.RootexecFlagOverride, "/bin/ksh -l")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		fmt.Println(err)
		return 2
	}

	return 0
}
