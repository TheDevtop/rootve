package main

import (
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/TheDevtop/ipcfs/go/ipcfs"
	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/TheDevtop/rootve/pkg/libve"
)

// Global lock and table
var (
	vtab libcsrv.VeTable
	lock sync.Mutex
)

// Global server (/srv/rootd)
var srv *net.UnixListener

func main() {
	var (
		err error
		mux *http.ServeMux
		mvc map[string]libve.VirtConfig
	)

	// Welcome message
	log.Println("Starting RootVE Server...")

	// Register server endpoint
	if srv, err = ipcfs.RegisterNetwork("rootd"); err != nil {
		log.Fatalln(err)
	}
	log.Println("Registered server endpoint")

	// Read /etc/rootve, initialize vtab
	if mvc, err = libve.ReadConfig(libve.ConfigPath); err != nil {
		log.Fatalln(err)
	}
	vtab = libcsrv.MakeTable(mvc)
	log.Println("Initialized table and configuration")

	// Autoboot enabled VE's
	autoboot()

	// Initialize multiplexer
	mux = http.NewServeMux()
	mux.HandleFunc(libcsrv.RouteStart, apiStart)
	mux.HandleFunc(libcsrv.RouteStop, apiStop)
	mux.HandleFunc(libcsrv.RouteListAll, apiListAll)
	mux.HandleFunc(libcsrv.RouteListOnline, apiListOnline)
	mux.HandleFunc(libcsrv.RoutePause, apiPause)
	mux.HandleFunc(libcsrv.RouteResume, apiResume)
	log.Println("Initialized multiplexer")

	// Setup signal listener
	go sigListen()
	log.Println("Initialized signal listener")

	// Serve WebAPI
	log.Println("Serving API")
	if err = http.Serve(srv, mux); err != nil {
		log.Println(err)
	}
}
