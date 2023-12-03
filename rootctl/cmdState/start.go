package cmdState

import (
	"bytes"
	"fmt"
	"os"

	"github.com/TheDevtop/rootve/pkg/jmap"
	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

const TagStart = "start"

func StartMain() {
	var (
		err    error
		buf    = new(bytes.Buffer)
		client = libcsrv.MakeClient()
	)

	if len(os.Args) < 2 {
		fmt.Println("Usage: start [name]")
		os.Exit(2)
	}

	if err = jmap.Mapto[[]byte]([]byte(os.Args[1]), buf); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if _, err = client.Post(libcsrv.MapProtocol(libcsrv.RouteListOnline), "", buf); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	os.Exit(0)
}
