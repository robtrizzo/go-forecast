package main

import (
	"flag"
	"log/slog"
	"os"
	"strings"

	"goforecast.robtrizzo/internal/weather"
)

type config struct {
	port int
	cors struct {
		trustedOrigins []string
	}
	weatherURL string
	// in a real project this would also include environment, rate limiter, and other top level configs
}

type application struct {
	config  config
	logger  *slog.Logger
	weather weather.Weather
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.weatherURL, "weather-url", "https://api.weather.gov", "weather api URL")
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	flag.Parse()

	weather := weather.NewWeather(cfg.weatherURL)

	// if this were an appliaction that needed a db, this is where we'd open a connection

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{config: cfg, logger: logger, weather: weather}

	err := app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
