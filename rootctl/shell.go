package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

const cmdShell = "shell"

// Run shell via rootexec
func runShell() error {
	cmd := exec.Command(libcsrv.RootexecPath, libcsrv.RootexecFlagName, os.Args[1], libcsrv.RootexecFlagOverride, "/bin/ksh -l")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func shellMain() int {
	var (
		err     error
		res     *http.Response
		resForm = new(libcsrv.Form[bool])
		client  = libcsrv.MakeClient()
		body    *strings.Reader
	)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [name]\n", cmdShell)
		return 2
	}

	if body, err = libcsrv.MakeJsonReader(libcsrv.FormMessage{
		Error: false,
		Data:  os.Args[1],
	}); err != nil {
		fmt.Println(err)
		return 2
	}

	if res, err = client.Post(libcsrv.MapProtocol(libcsrv.RouteOnline), "", body); err != nil {
		fmt.Println(err)
		return 2
	}
	if err = libcsrv.ReadJson(res.Body, resForm); err != nil {
		fmt.Println(err)
		return 2
	}

	if resForm.Error || !resForm.Data {
		fmt.Printf("Could not login to %s\n", os.Args[1])
		return 2
	}

	if err = runShell(); err != nil {
		fmt.Println(err)
		return 2
	}

	return 0
}
