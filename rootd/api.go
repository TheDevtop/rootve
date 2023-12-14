package main

import (
	"log"
	"net/http"

	"github.com/TheDevtop/rootve/pkg/libcsrv"
	"github.com/TheDevtop/rootve/pkg/librex"
)

// Start a named Virtual Environment
func apiStart(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		rex      *librex.Rex
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

	// Begin critical section
	globalRexMap.Lock.Lock()
	if rex = globalRexMap.Map[nameForm.Data]; rex != nil {
		err = rex.Start()
		globalRexMap.Map[nameForm.Data] = rex
	}
	globalRexMap.Lock.Unlock()
	// End critical section

	// Send a response message
	if err != nil {
		libcsrv.WriteJson(w, libcsrv.FormMessage{
			Error: true,
			Data:  err.Error(),
		})
		log.Println(err)
		return
	}

	libcsrv.WriteJson(w, libcsrv.FormMessage{
		Error: false,
		Data:  "",
	})
}

// Stop a named Virtual Environment
func apiStop(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		rex      *librex.Rex
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

	// Begin critical section
	globalRexMap.Lock.Lock()
	if rex = globalRexMap.Map[nameForm.Data]; rex != nil {
		err = rex.Stop()
		globalRexMap.Map[nameForm.Data] = rex
	}
	globalRexMap.Lock.Unlock()
	// End critical section

	// Send a response message
	if err != nil {
		libcsrv.WriteJson(w, libcsrv.FormMessage{
			Error: true,
			Data:  err.Error(),
		})
		log.Println(err)
		return
	}

	libcsrv.WriteJson(w, libcsrv.FormMessage{
		Error: false,
		Data:  "",
	})
}

// List all Virtual Environments
func apiListAll(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		form = new(libcsrv.Form[[]libcsrv.FormVeList])
	)

	// Begin critical section
	globalRexMap.Lock.Lock()
	for key, rex := range globalRexMap.Map {
		if rex != nil {
			form.Data = append(form.Data, libcsrv.FormVeList{
				Name:    key,
				State:   librex.StateToLabel(rex.State),
				Path:    rex.Config.Root,
				Command: rex.Config.CommandPath,
			})
		}
	}
	globalRexMap.Lock.Unlock()
	// End critical section

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

	// Begin critical section
	globalRexMap.Lock.Lock()
	for key, rex := range globalRexMap.Map {
		if rex != nil {
			if rex.State == librex.StateOn {
				form.Data = append(form.Data, libcsrv.FormVeList{
					Name:    key,
					State:   librex.StateToLabel(rex.State),
					Path:    rex.Config.Root,
					Command: rex.Config.CommandPath,
				})
			}
		}
	}
	globalRexMap.Lock.Unlock()
	// End critical section

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
