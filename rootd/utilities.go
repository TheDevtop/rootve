package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"golang.org/x/sys/unix"
)

// Listen for exit signal, shutdown server
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

	close(ch)
	os.Exit(0)
}

// Autoboot VE's where autoboot=true
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
