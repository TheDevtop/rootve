package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

const cmdStop = "stop"

func stopMain() int {
	var (
		err     error
		res     *http.Response
		resForm = new(libcsrv.FormMessage)
		client  = libcsrv.MakeClient()
		body    *strings.Reader
	)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [name]\n", cmdStop)
		return 2
	}

	if body, err = libcsrv.MakeJsonReader(libcsrv.FormMessage{
		Error: false,
		Data:  os.Args[1],
	}); err != nil {
		fmt.Println(err)
		return 2
	}

	if res, err = client.Post(libcsrv.MapProtocol(libcsrv.RouteStop), "", body); err != nil {
		fmt.Println(err)
		return 2
	}
	if err = libcsrv.ReadJson(res.Body, resForm); err != nil {
		fmt.Println(err)
		return 2
	}

	if resForm.Error {
		fmt.Printf("Could not stop %s\n", os.Args[1])
		return 2
	}

	fmt.Printf("Stopped %s\n", os.Args[1])
	return 0
}
