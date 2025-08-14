package weather

import (
	"errors"
	"fmt"
	"net/http"

	"goforecast.robtrizzo/internal/data"
)

var (
	ErrGridPointDataUnavailable = errors.New("data unavailable for requested gridpoint")
	ErrNoForecast               = errors.New("no forecast data available for requested gridpoint")
)

// in a real project, these may be optional runtime flags
const (
	COLD = 50
	HOT  = 80
)

type Weather struct {
	url string
}

func NewWeather(url string) Weather {
	return Weather{
		url: url,
	}
}

func (w Weather) GetForecastURLFromCoordinates(l *data.Coordinate) (string, error) {
	endpoint := fmt.Sprintf("%s/points/%v,%v", w.url, l.Latitude, l.Longitude)
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 404 {
		return "", ErrGridPointDataUnavailable
	}

	// in a real project this would also include all of the other
	// keys from the response
	var forecastFromCoordinatesResponse struct {
		Properties struct {
			Forecast string `json:"forecast"`
		} `json:"properties"`
	}
	err = w.readJSON(resp.Body, &forecastFromCoordinatesResponse)
	if err != nil {
		return "", err
	}
	return forecastFromCoordinatesResponse.Properties.Forecast, nil
}

func (w Weather) GetForecast(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// in a real project this would also include all of the other
	// keys from the response
	var forecastResponse struct {
		Properties struct {
			Periods []data.Period `json:"periods"`
		} `json:"properties"`
	}
	err = w.readJSON(resp.Body, &forecastResponse)
	if err != nil {
		return "", err
	}
	if len(forecastResponse.Properties.Periods) == 0 {
		return "", ErrNoForecast
	}

	periodToday := forecastResponse.Properties.Periods[0]
	var tempSummary string
	switch {
	case periodToday.Temperature <= COLD:
		tempSummary = "cold"
	case periodToday.Temperature >= HOT:
		tempSummary = "hot"
	default:
		tempSummary = "moderate"
	}

	return fmt.Sprintf("Today's forecast is %s and %s.", periodToday.ShortForecast, tempSummary), nil
}
