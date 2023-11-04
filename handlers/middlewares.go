package handlers

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (app *Application) Middleware(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		handler(w, r, ps)
	}
}
