package cmdList

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/TheDevtop/rootve/pkg/libcsrv/models"
	"github.com/da0x/golang/olog"
)

const TagLs = "ls"

func LsMain() {
	var (
		err    error
		resp   *http.Response
		client = libcsrv.MakeClient()
		tab    libcsrv.VeTable
		list   models.VeList
	)

	if resp, err = client.Get(libcsrv.MapProtocol(libcsrv.RouteListAll)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if tab, err = libcsrv.ReadTable(resp.Body); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for k, v := range tab {
		list = append(list, models.VeEntry{
			Name:    k,
			Path:    v.Config.Root,
			State:   v.State,
			Command: v.Config.CommandPath,
		})
	}

	olog.Print(list)
	os.Exit(0)
}
