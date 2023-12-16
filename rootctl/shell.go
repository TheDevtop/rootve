package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/TheDevtop/rootve/pkg/librex"
)

const cmdShell = "shell"

// Run shell via rootexec
func runShell() error {
	cmd := exec.Command(librex.RootexecPath, librex.RootexecFlagName, os.Args[1], librex.RootexecFlagOverride, "/bin/ksh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func shellMain() int {
	var err error

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [name]\n", cmdShell)
		return 2
	}

	if err = runShell(); err != nil {
		fmt.Println(err)
		return 2
	}

	return 0
}
