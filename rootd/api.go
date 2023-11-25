package main

import (
	"log"
	"net/http"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

func apiStart(w http.ResponseWriter, r *http.Request) {
	var (
		name  = r.Header.Get(libcsrv.HdrName)
		entry *libcsrv.VeEntry
		err   error
	)

	// Critical section
	lock.Lock()
	entry = vtab[name]
	lock.Unlock()

	if entry == nil {
		log.Println("Requested entry is nil pointer")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = entry.Start(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Started %s\n", name)
}

func apiStop(w http.ResponseWriter, r *http.Request) {
	var (
		name  = r.Header.Get(libcsrv.HdrName)
		entry *libcsrv.VeEntry
		err   error
	)

	// Critical section
	lock.Lock()
	entry = vtab[name]
	lock.Unlock()

	if entry == nil {
		log.Println("Requested entry is nil pointer")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = entry.Stop(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Stopped %s\n", name)
}

func apiListAll(w http.ResponseWriter, r *http.Request) {

}

func apiListOnline(w http.ResponseWriter, r *http.Request) {

}

func apiPause(w http.ResponseWriter, r *http.Request) {
	var (
		name  = r.Header.Get(libcsrv.HdrName)
		entry *libcsrv.VeEntry
		err   error
	)

	// Critical section
	lock.Lock()
	entry = vtab[name]
	lock.Unlock()

	if entry == nil {
		log.Println("Requested entry is nil pointer")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = entry.Pause(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Paused %s\n", name)
}

func apiResume(w http.ResponseWriter, r *http.Request) {
	var (
		name  = r.Header.Get(libcsrv.HdrName)
		entry *libcsrv.VeEntry
		err   error
	)

	// Critical section
	lock.Lock()
	entry = vtab[name]
	lock.Unlock()

	if entry == nil {
		log.Println("Requested entry is nil pointer")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = entry.Resume(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Resumed %s\n", name)
}
