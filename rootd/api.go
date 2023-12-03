package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TheDevtop/rootve/pkg/jmap"
	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

func apiStart(w http.ResponseWriter, r *http.Request) {
	var (
		name  string
		entry *libcsrv.VeEntry
		err   error
	)

	if err = jmap.Mapfrom[string](r.Body, &name); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Critical section
	lock.Lock()
	entry = vtab[name]
	lock.Unlock()

	if entry == nil {
		log.Printf("Requested entry (%s) is nil pointer\n", name)
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
		name  string
		entry *libcsrv.VeEntry
		err   error
	)

	if err = jmap.Mapfrom[string](r.Body, &name); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Critical section
	lock.Lock()
	entry = vtab[name]
	lock.Unlock()

	if entry == nil {
		log.Printf("Requested entry (%s) is nil pointer\n", name)
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
	var (
		err error
		buf []byte
	)

	// Critical section
	lock.Lock()
	buf, err = json.Marshal(vtab)
	lock.Unlock()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(buf)
}

func apiListOnline(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		buf  []byte
		otab = make(libcsrv.VeTable)
	)

	// Critical section
	lock.Lock()
	for k, v := range vtab {
		if v.State == libcsrv.StateOn {
			otab[k] = v
		}
	}
	lock.Unlock()

	if buf, err = json.Marshal(otab); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(buf)
}

func apiPause(w http.ResponseWriter, r *http.Request) {
	var (
		name  string
		entry *libcsrv.VeEntry
		err   error
	)

	if err = jmap.Mapfrom[string](r.Body, &name); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Critical section
	lock.Lock()
	entry = vtab[name]
	lock.Unlock()

	if entry == nil {
		log.Printf("Requested entry (%s) is nil pointer\n", name)
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
		name  string
		entry *libcsrv.VeEntry
		err   error
	)

	if err = jmap.Mapfrom[string](r.Body, &name); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Critical section
	lock.Lock()
	entry = vtab[name]
	lock.Unlock()

	if entry == nil {
		log.Printf("Requested entry (%s) is nil pointer\n", name)
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
