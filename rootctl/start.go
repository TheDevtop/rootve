package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

const cmdStart = "start"

// Request the server to start a virtual environment
func startMain() int {
	var (
		err     error
		res     *http.Response
		resForm = new(libcsrv.FormMessage)
		client  = libcsrv.MakeClient()
		body    *strings.Reader
	)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [name]\n", cmdStart)
		return 2
	}

	if body, err = libcsrv.MakeJsonReader(libcsrv.FormMessage{
		Error: false,
		Data:  os.Args[1],
	}); err != nil {
		fmt.Println(err)
		return 2
	}

	if res, err = client.Post(libcsrv.MapProtocol(libcsrv.RouteStart), "", body); err != nil {
		fmt.Println(err)
		return 2
	}
	if err = libcsrv.ReadJson(res.Body, resForm); err != nil {
		fmt.Println(err)
		return 2
	}

	if resForm.Error {
		fmt.Printf("Could not start %s\n", os.Args[1])
		return 2
	}

	fmt.Printf("Started %s\n", os.Args[1])
	return 0
}
