package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"goforecast.robtrizzo/internal/validator"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readFloat(qs url.Values, key string, defaultValue float64, v *validator.Validator) float64 {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		v.AddError(key, "must be a decimal value")
		return float64(defaultValue)
	}

	return f
}
