package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

const cmdPause = "pause"

func pauseMain() int {
	var (
		err     error
		res     *http.Response
		resForm = new(libcsrv.FormMessage)
		client  = libcsrv.MakeClient()
		body    *strings.Reader
	)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [name]\n", cmdPause)
		return 2
	}

	if body, err = libcsrv.MakeJsonReader(libcsrv.FormMessage{
		Error: false,
		Data:  os.Args[1],
	}); err != nil {
		fmt.Println(err)
		return 2
	}

	if res, err = client.Post(libcsrv.MapProtocol(libcsrv.RoutePause), "", body); err != nil {
		fmt.Println(err)
		return 2
	}
	if err = libcsrv.ReadJson(res.Body, resForm); err != nil {
		fmt.Println(err)
		return 2
	}

	if resForm.Error {
		fmt.Printf("Could not pause %s\n", os.Args[1])
		return 2
	}

	fmt.Printf("Paused %s\n", os.Args[1])
	return 0
}
