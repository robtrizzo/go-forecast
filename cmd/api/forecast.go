package main

import (
	"errors"
	"fmt"
	"net/http"

	"goforecast.robtrizzo/internal/data"
	"goforecast.robtrizzo/internal/validator"
	"goforecast.robtrizzo/internal/weather"
)

func (app *application) forecastHandler(w http.ResponseWriter, r *http.Request) {
	var input data.Coordinate

	v := validator.New()

	qs := r.URL.Query()

	// default values are for Detroit, but this would be a good place to read the
	// requester's location if available, then use that as the default value
	input.Latitude = app.readFloat(qs, "lat", 42.3297, v)
	input.Longitude = app.readFloat(qs, "lon", -83.0425, v)

	if data.ValidateCoordinate(v, input); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	gp, err := app.weather.GetForecastURLFromCoordinates(&input)
	if err != nil {
		switch {
		case errors.Is(err, weather.ErrGridPointDataUnavailable):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	forecast, err := app.weather.GetForecast(gp)
	if err != nil {
		switch {
		case errors.Is(err, weather.ErrNoForecast):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%s\n", forecast)
}
