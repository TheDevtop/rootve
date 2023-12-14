package main

import (
	"log"
	"net"
	"net/http"

	"github.com/TheDevtop/ipcfs/go/ipcfs"
	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/TheDevtop/rootve/pkg/librex"
	"github.com/TheDevtop/rootve/pkg/libve"
)

// Global data
var (
	store librex.RexMap     // Stores the rootexec instances
	srv   *net.UnixListener // Server
)

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

	// Read /etc/rootve, initialize RexMap
	if mvc, err = libve.ReadConfig(libve.ConfigPath); err != nil {
		log.Fatalln(err)
	}
	store = ConfigToRexMap(mvc)
	log.Println("Initialized global structures")

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
	mux.HandleFunc(libcsrv.RouteOnline, apiOnline)
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
