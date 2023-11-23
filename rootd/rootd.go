package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/TheDevtop/rootve/pkg/libve"
	"golang.org/x/sys/unix"
)

// Global lock and table
var (
	vtab libcsrv.VeTable
	lock sync.Mutex
)

func sigListen() {
	ch := make(chan os.Signal, 1)

	// Wait for the signal
	signal.Notify(ch, os.Interrupt, unix.SIGTERM)
	<-ch

	// Stop the environments
	lock.Lock()
	for key, val := range vtab {
		val.Exec.Cancel()
		val.State = libcsrv.StateOff
		log.Printf("Stopped %s\n", key)
	}
	lock.Unlock()

	os.Exit(0)
}

func autoboot() {
	var err error

	lock.Lock()
	for key, val := range vtab {
		if val.Config.Autoboot {
			if err = val.Exec.Start(); err != nil {
				log.Printf("Could not autoboot %s: %s\n", key, err)
			} else {
				val.State = libcsrv.StateOn
				log.Printf("Autobooted %s\n", key)
			}
		}
	}
	lock.Unlock()
}

func main() {
	var (
		err    error
		socket net.Listener
		mux    *http.ServeMux
		mvc    map[string]libve.VirtConfig
	)

	// Welcome message
	log.Println("Starting RootVE Server...")

	// Allocate socket
	if socket, err = net.Listen("unix", libcsrv.SocketPath); err != nil {
		log.Panicln(err)
	}
	log.Printf("Socket endpoint: %s\n", libcsrv.SocketPath)

	// Read /etc/rootve, initialize vtab
	if mvc, err = libve.ReadConfig(libve.DefaultPath); err != nil {
		log.Fatalln(err)
	}
	vtab = libcsrv.MakeTable(mvc)
	log.Println("Initialized table and configuration")

	// Autoboot enabled VE's
	autoboot()

	// Setup HTTP multiplexer
	mux.HandleFunc(libcsrv.RouteStart, apiStart)
	mux.HandleFunc(libcsrv.RouteStop, apiStop)
	mux.HandleFunc(libcsrv.RouteListAll, apiListAll)
	mux.HandleFunc(libcsrv.RouteListOnline, apiListOnline)
	mux.HandleFunc(libcsrv.RoutePause, apiPause)
	mux.HandleFunc(libcsrv.RouteResume, apiResume)

	// Setup signal listener
	go sigListen()
	log.Println("Initialized signal listener")

	// Serve WebAPI
	log.Println("Serving API")
	if err = http.Serve(socket, mux); err != nil {
		log.Println(err)
	}
}
