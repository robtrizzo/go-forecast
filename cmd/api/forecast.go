package main

import (
	"fmt"
	"net/http"

	"goforecast.robtrizzo/internal/data"
	"goforecast.robtrizzo/internal/validator"
)

func (app *application) forecastHandler(w http.ResponseWriter, r *http.Request) {
	var input data.Location

	v := validator.New()

	qs := r.URL.Query()

	// default values are for Detroit, but this would be a good place to read the
	// requester's location if available, then use that as the default value
	input.Latitude = app.readFloat(qs, "lat", 42.3297, v)
	input.Longitude = app.readFloat(qs, "lon", 83.0425, v)

	if data.ValidateLocation(v, input); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}
