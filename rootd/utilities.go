package main

import (
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
	store.LockPtr.Lock()
	for key, vmp := range store.MapPtr {
		if vmp != nil {
			if vmp.proc != nil {
				if vmp.Switch(libcsrv.StateOff) == nil {
					log.Printf("Stopped %s\n", key)
				}
			}
		}
	}
	store.LockPtr.Lock()
}

// Autoboot VE's where autoboot=true
func autoboot() {
	var err error
	for key, vmp := range vmap {
		if vmp.config.Autoboot && vmp.proc != nil {
			if err = vmp.Switch(libcsrv.StateOn); err != nil {
				log.Printf("Could not autoboot %s: %s\n", key, err)
			} else {
				log.Printf("Autobooted %s\n", key)
			}
		}
	}
}

// Convert configmap to rexmap
func ConfigToRexMap(mvc map[string]libve.VirtConfig) librex.RexMap {
	rm := librex.MakeRexMap(len(mvc))
	for name, vc := range mvc {
		rm.Store(name, librex.NewRex(name, vc))
	}
	return rm
}
