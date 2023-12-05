package main

import (
	"log"
	"net/http"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

// Start a named Virtual Environment
func apiStart(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		vmp      *vmach
		nameForm = new(libcsrv.FormMessage)
	)

	// Read the name from the form
	if err = libcsrv.ReadJson(r.Body, nameForm); err != nil {
		log.Println(err)
		libcsrv.WriteJson(w, libcsrv.FormMessage{
			Error: true,
			Data:  err.Error(),
		})
		return
	}

	// Find the vmp, critical section
	lock.Lock()
	if vmp = vmap[nameForm.Data]; vmp != nil {
		log.Printf("%s: %s\n", errVmapEntry, nameForm.Data)
		libcsrv.WriteJson(w, libcsrv.FormMessage{
			Error: true,
			Data:  errVmapEntry.Error(),
		})
		return
	}
	lock.Unlock()

	// Attempt to start the vmp
	if err = vmp.Switch(libcsrv.StateOn); err != nil {
		log.Println(err)
		libcsrv.WriteJson(w, libcsrv.FormMessage{
			Error: true,
			Data:  err.Error(),
		})
		return
	}

	// Send a response message
	if err = libcsrv.WriteJson(w, libcsrv.FormMessage{
		Error: false,
		Data:  "",
	}); err != nil {
		log.Println(err)
	}
}

// Start a named Virtual Environment
func apiStop(w http.ResponseWriter, r *http.Request) {

}

// List all Virtual Environments
func apiListAll(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		form = new(libcsrv.Form[[]libcsrv.FormVeList])
	)

	// Critical section
	lock.Lock()
	for key, vmp := range vmap {
		if vmp != nil {
			form.Data = append(form.Data, libcsrv.FormVeList{
				Name:    key,
				State:   libcsrv.Slabel(vmp.state),
				Path:    vmp.config.Root,
				Command: vmp.config.CommandPath,
			})
		}
	}
	lock.Unlock()

	if err = libcsrv.WriteJson(w, *form); err != nil {
		log.Println(err)
	}
}

// List online Virtual Environments
func apiListOnline(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		form = new(libcsrv.Form[[]libcsrv.FormVeList])
	)

	// Critical section
	lock.Lock()
	for key, vmp := range vmap {
		if vmp != nil && vmp.state == libcsrv.StateOn {
			form.Data = append(form.Data, libcsrv.FormVeList{
				Name:    key,
				State:   libcsrv.Slabel(vmp.state),
				Path:    vmp.config.Root,
				Command: vmp.config.CommandPath,
			})
		}
	}
	lock.Unlock()

	if err = libcsrv.WriteJson(w, *form); err != nil {
		log.Println(err)
	}
}

// Pause a named Virtual Environment
func apiPause(w http.ResponseWriter, r *http.Request) {

}

// Resume a named Virtual Environment
func apiResume(w http.ResponseWriter, r *http.Request) {

}

// Assert if virtual environment is online
func apiOnline(w http.ResponseWriter, r *http.Request) {
	var (
		nameForm   = new(libcsrv.FormMessage)
		onlineForm = new(libcsrv.Form[bool])
		err        error
		vmp        *vmach
	)

	if err = libcsrv.ReadJson(r.Body, nameForm); err != nil {
		log.Println(err)
		libcsrv.WriteJson(w, libcsrv.FormMessage{
			Error: true,
			Data:  err.Error(),
		})
		return
	}

	// Critical section
	lock.Lock()
	if vmp = vmap[nameForm.Data]; vmp != nil && vmp.state == libcsrv.StateOn {
		onlineForm.Data = true
	} else {
		onlineForm.Data = false
	}
	lock.Unlock()

	if err = libcsrv.WriteJson(w, *onlineForm); err != nil {
		log.Println(err)
	}
}
