package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RemoteAddr, r.RequestURI)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/api/add", app.addRecord)
	router.HandlerFunc(http.MethodGet, "/api/get", app.getAllRecord)
	router.HandlerFunc(http.MethodGet, "/api/getMonth", app.getRecordInMonth)
	return app.enableCORS(router)
}
