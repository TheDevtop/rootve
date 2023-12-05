package cmdShell

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

const TagShell = "shell"

// Run shell via rootexec
func runShell() error {
	cmd := exec.Command(libcsrv.RootexecPath, libcsrv.RootexecFlagName, os.Args[1], libcsrv.RootexecFlagOverride, "/bin/ksh -l")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ShellMain() {
	var (
		err     error
		res     *http.Response
		resForm = new(libcsrv.Form[bool])
		client  = libcsrv.MakeClient()
		body    *strings.Reader
	)

	if len(os.Args) < 2 {
		fmt.Println("Usage: shell [name]")
		os.Exit(2)
	}

	if body, err = libcsrv.MakeJsonReader(libcsrv.FormMessage{
		Error: false,
		Data:  os.Args[1],
	}); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if res, err = client.Post(libcsrv.MapProtocol(libcsrv.RouteOnline), "", body); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if err = libcsrv.ReadJson(res.Body, resForm); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if resForm.Error || !resForm.Data {
		fmt.Printf("Could not determine the state of %s\n", os.Args[1])
		os.Exit(2)
	}

	if err = runShell(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	os.Exit(0)
}
