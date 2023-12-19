package main

import (
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/TheDevtop/ipcfs/go/ipcfs"
	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/TheDevtop/rootve/pkg/librex"
	"github.com/TheDevtop/rootve/pkg/libve"
	"golang.org/x/sys/unix"
)

// Listen for exit signal, shutdown server
func sigListen() {
	ch := make(chan os.Signal, 1)

	// Wait for the signal
	signal.Notify(ch, unix.SIGINT, unix.SIGTERM, unix.SIGSTOP)
	<-ch

	// Deregister server endpoint
	ipcfs.DeregisterNetwork("rootd", globalServer)

	// Stop the environments
	autohalt()

	close(ch)
	os.Exit(0)
}

// Halt instances
func autohalt() {
	globalRexMap.Lock.Lock()
	for key, rex := range globalRexMap.Map {
		if rex != nil {
			if err := rex.Stop(); err != nil {
				logError("Autohalt", err)
			}
		}
		log.Printf("Stopped: %s\n", key)
	}
	globalRexMap.Lock.Unlock()
}

// Autoboot instances where autoboot=true
func autoboot() {
	globalRexMap.Lock.Lock()
	for key, rex := range globalRexMap.Map {
		if rex.Config.Autoboot {
			if err := rex.Start(); err != nil {
				logError("Autoboot", err)
			}
			log.Printf("Started: %s\n", key)
		}
	}
	globalRexMap.Lock.Unlock()
}

// Send response with error message
func responseError(w io.Writer, err error) {
	libcsrv.WriteJson(w, libcsrv.FormMessage{
		Error: true,
		Data:  err.Error(),
	})
}

// Log error message
func logError(path string, err error) {
	log.Printf("%s: %s\n", path, err)
}

// Convert configmap to rexmap
func ConfigToRexMap(mvc map[string]libve.VirtConfig) librex.RexMap {
	rm := librex.MakeRexMap(len(mvc))
	for name, vc := range mvc {
		rm.Map[name] = librex.NewRex(name, vc)
	}
	return rm
}
