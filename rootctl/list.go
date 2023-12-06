package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/da0x/golang/olog"
)

const (
	cmdLs = "ls"
	cmdPs = "ps"
)

func lsMain() int {
	var (
		err     error
		res     *http.Response
		resForm = new(libcsrv.Form[[]libcsrv.FormVeList])
		client  = libcsrv.MakeClient()
	)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [name]\n", cmdLs)
		return 2
	}

	if res, err = client.Get(libcsrv.MapProtocol(libcsrv.RouteListAll)); err != nil {
		fmt.Println(err)
		return 2
	}

	if err = libcsrv.ReadJson(res.Body, resForm); err != nil {
		fmt.Println(err)
		return 2
	}

	olog.Print(resForm.Data)
	return 0
}

func psMain() int {
	var (
		err     error
		res     *http.Response
		resForm = new(libcsrv.Form[[]libcsrv.FormVeList])
		client  = libcsrv.MakeClient()
	)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [name]\n", cmdPs)
		return 2
	}

	if res, err = client.Get(libcsrv.MapProtocol(libcsrv.RouteListOnline)); err != nil {
		fmt.Println(err)
		return 2
	}

	if err = libcsrv.ReadJson(res.Body, resForm); err != nil {
		fmt.Println(err)
		return 2
	}

	olog.Print(resForm.Data)
	return 0
}