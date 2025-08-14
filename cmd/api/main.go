package main

import (
	"flag"
	"log/slog"
	"os"
	"strings"
)

type config struct {
	port int
	cors struct {
		trustedOrigins []string
	}
	// in a real project this would also include environment, rate limiter, and other top level configs
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	flag.Parse()

	// if this were an appliaction that needed a db, this is where we'd open a connection

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{config: cfg, logger: logger}

	err := app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
