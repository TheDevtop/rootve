package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/TheDevtop/ipcfs/go/ipcfs"
	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/TheDevtop/rootve/pkg/libve"
	"golang.org/x/sys/unix"
)

// Listen for exit signal, shutdown server
func sigListen() {
	ch := make(chan os.Signal, 1)

	// Wait for the signal
	signal.Notify(ch, os.Interrupt, unix.SIGTERM)
	<-ch

	// Deregister server endpoint
	ipcfs.DeregisterNetwork("rootd", srv)

	// Stop the environments
	autostop()

	close(ch)
	os.Exit(0)
}

// Autostop the VE's
func autostop() {
	lock.Lock()
	for key, vmp := range vmap {
		if vmp != nil {
			if vmp.proc != nil {
				if vmp.Switch(libcsrv.StateOff) == nil {
					log.Printf("Stopped %s\n", key)
				}
			}
		}
	}
	lock.Unlock()
}

// Autoboot VE's where autoboot=true
func autoboot() {
	var err error

	lock.Lock()
	for key, vmp := range vmap {
		if vmp.config.Autoboot && vmp.proc != nil {
			if err = vmp.Switch(libcsrv.StateOn); err != nil {
				log.Printf("Could not autoboot %s: %s\n", key, err)
			} else {
				log.Printf("Autobooted %s\n", key)
			}
		}
	}
	lock.Unlock()
}

// Allocate and initialize a "Virtual Machine Map"
func makeVmap(mvc map[string]libve.VirtConfig) map[string]*vmach {
	newMap := make(map[string]*vmach, len(mvc))
	for name, vc := range mvc {
		newMap[name] = newVmach(name, vc)
	}
	return newMap
}
