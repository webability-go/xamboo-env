// package box is an engine for Xamboo.
// The engine will just returns a string with some data in a box in HTML format
package main

import (
	"net/http"
)

var Component = MyHandler{}

type MyHandler struct{}

func (hndl *MyHandler) Start() {
}

func (hndl *MyHandler) NeedHandler() bool {
	return true
}

func (hndl *MyHandler) Handler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		http.Error(w, "404 Error de MyHandler", http.StatusNotFound)
		return
	}
}
