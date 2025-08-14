package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// if this application needed authentication and/or authorization,
	// that middleware could be used at the route level, for example wrapping the handler
	router.HandlerFunc(http.MethodGet, "/v1/forecast", app.forecastHandler)

	// CORS is overkill for this app, but I wanted to demonstrate at least one middleware
	// return app.enableCORS(router)
	return router
}
