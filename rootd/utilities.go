package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/TheDevtop/ipcfs/go/ipcfs"
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
	ipcfs.DeregisterNetwork("rootd", globalServer)

	// Stop the environments
	autostop()

	close(ch)
	os.Exit(0)
}

// Autostop the VE's
func autostop() {
	globalRexMap.Lock.Lock()
	for key, rex := range globalRexMap.Map {
		if err := rex.Stop(); err != nil {
			log.Println(err)
		}
		log.Printf("Stopped: %s\n", key)
	}
	globalRexMap.Lock.Unlock()
}

// Autoboot VE's where autoboot=true
func autoboot() {
	globalRexMap.Lock.Lock()
	for key, rex := range globalRexMap.Map {
		if rex.Config.Autoboot {
			if err := rex.Start(); err != nil {
				log.Println(err)
			}
			log.Printf("Started: %s\n", key)
		}
	}
	globalRexMap.Lock.Unlock()
}

// Convert configmap to rexmap
func ConfigToRexMap(mvc map[string]libve.VirtConfig) librex.RexMap {
	rm := librex.MakeRexMap(len(mvc))
	for name, vc := range mvc {
		rm.Map[name] = librex.NewRex(name, vc)
	}
	return rm
}
