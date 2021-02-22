// package box is an engine for Xamboo.
// The engine will just returns a string with some data in a box in HTML format
package main

import (
	"fmt"
	"net/http"

	"github.com/webability-go/xamboo/components/host"
	"github.com/webability-go/xamboo/config"
)

var Component = MyHandler{}

type MyHandler struct{}
type MyHandlerConfig struct {
	Error404    bool
	Maintenance bool
	Page        string
}

func (hndl *MyHandler) Start() {
}

func (hndl *MyHandler) StartHost(host *config.Host) {

	config := MyHandlerConfig{}

	mycomponent := host.Components["myhandler"]
	cfg := mycomponent.Params

	config.Error404, _ = cfg["404"].(bool)
	config.Maintenance, _ = cfg["maintenance"].(bool)
	config.Page, _ = cfg["page"].(string)

	mycomponent.Config = config
}

func (hndl *MyHandler) NeedHandler() bool {
	return true
}

func (hndl *MyHandler) Handler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get Config Parameters

		hw, ok := w.(host.HostWriter)
		if !ok {
			fmt.Println("FileServer component: ERROR DETECTED: the writer is not a HostWriter or the Listener/Host is not set (and that should not happen)", r, w)
			http.Error(w, "FileServer component: Writer error", http.StatusInternalServerError)
			return
		}
		host := hw.GetHost()
		if host == nil {
			fmt.Println("FileServer component: ERROR DETECTED: there is no HOST (and that should not happen)", r, w)
			http.Error(w, "FileServer component: Writer error (see logs for more info)", http.StatusInternalServerError)
			return
		}
		mycomponent := host.Components["myhandler"]
		if mycomponent != nil && mycomponent.Enabled {
			myconfig, _ := mycomponent.Config.(MyHandlerConfig)

			if myconfig.Error404 {
				http.Error(w, "404 Error de MyHandler", http.StatusNotFound)
				return
			}

			if myconfig.Maintenance {
				http.Error(w, "200 MANTENIMIENTO", http.StatusOK)
				return
			}
		}

		handler.ServeHTTP(w, r)
	}
}
