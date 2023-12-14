package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

const cmdRemove = "rm"

func removeMain() int {
	var (
		err     error
		res     *http.Response
		resForm = new(libcsrv.FormMessage)
		client  = libcsrv.MakeClient()
		body    *strings.Reader
	)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [name]\n", cmdRemove)
		return 2
	}

	if body, err = libcsrv.MakeJsonReader(libcsrv.FormMessage{
		Error: false,
		Data:  os.Args[1],
	}); err != nil {
		fmt.Println(err)
		return 2
	}

	if res, err = client.Post(libcsrv.MapProtocol(libcsrv.RouteRemove), "", body); err != nil {
		fmt.Println(err)
		return 2
	}
	if err = libcsrv.ReadJson(res.Body, resForm); err != nil {
		fmt.Println(err)
		return 2
	}

	if resForm.Error {
		fmt.Printf("Could not remove %s\n", os.Args[1])
		return 2
	}

	fmt.Printf("Removed %s\n", os.Args[1])
	return 0
}
