package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/add", app.addRecord)
	router.HandlerFunc(http.MethodGet, "/v1/get", app.getAllRecord)
	router.HandlerFunc(http.MethodGet, "/v1/getMonth", app.getRecordInMonth)
	return router
}
