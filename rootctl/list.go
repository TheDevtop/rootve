package main

import (
	"fmt"
	"net/http"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/jedib0t/go-pretty/v6/table"
)

const (
	cmdLs = "ls"
	cmdPs = "ps"
)

// Request the server to return a list with all the VE
func lsMain() int {
	var (
		err     error
		res     *http.Response
		resForm = new(libcsrv.Form[[]libcsrv.FormVeList])
		client  = libcsrv.MakeClient()
		tw      = table.NewWriter()
	)

	if res, err = client.Get(libcsrv.MapProtocol(libcsrv.RouteListAll)); err != nil {
		fmt.Println(err)
		return 2
	}

	if err = libcsrv.ReadJson(res.Body, resForm); err != nil {
		fmt.Println(err)
		return 2
	}

	tw.AppendHeader(table.Row{"Name", "State", "Interface", "Address", "Command"})
	for _, entry := range resForm.Data {
		tw.AppendRow(table.Row{
			entry.Name,
			entry.State,
			entry.Interface,
			entry.Address,
			entry.Command,
		})
	}

	fmt.Println(tw.Render())
	return 0
}

// Request the server to return a list with the active VE
func psMain() int {
	var (
		err     error
		res     *http.Response
		resForm = new(libcsrv.Form[[]libcsrv.FormVeList])
		client  = libcsrv.MakeClient()
		tw      = table.NewWriter()
	)

	if res, err = client.Get(libcsrv.MapProtocol(libcsrv.RouteListOnline)); err != nil {
		fmt.Println(err)
		return 2
	}

	if err = libcsrv.ReadJson(res.Body, resForm); err != nil {
		fmt.Println(err)
		return 2
	}

	tw.AppendHeader(table.Row{"Name", "State", "Interface", "Address", "Command"})
	for _, entry := range resForm.Data {
		tw.AppendRow(table.Row{
			entry.Name,
			entry.State,
			entry.Interface,
			entry.Address,
			entry.Command,
		})
	}

	fmt.Println(tw.Render())
	return 0
}
