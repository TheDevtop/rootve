package main

import (
	"log"
	"net/http"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
)

// Start a named Virtual Environment
func apiStart(w http.ResponseWriter, r *http.Request) {

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
	for key, val := range vmap {
		if val != nil {
			form.Data = append(form.Data, libcsrv.FormVeList{
				Name:    key,
				State:   libcsrv.Slabel(val.state),
				Path:    val.config.Root,
				Command: val.config.CommandPath,
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
	for key, val := range vmap {
		if val != nil && val.state == libcsrv.StateOn {
			form.Data = append(form.Data, libcsrv.FormVeList{
				Name:    key,
				State:   libcsrv.Slabel(val.state),
				Path:    val.config.Root,
				Command: val.config.CommandPath,
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
