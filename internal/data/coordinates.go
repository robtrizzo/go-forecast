package data

import "goforecast.robtrizzo/internal/validator"

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func ValidateCoordinate(v *validator.Validator, l Coordinate) {
	v.Check(l.Latitude < 90, "lat", "must be between -90 and 90")
	v.Check(l.Latitude > -90, "lat", "must be between -90 and 90")
	v.Check(l.Longitude < 180, "lon", "must be between -180 and 180")
	v.Check(l.Longitude > -180, "lon", "must be between -180 and 180")
}
