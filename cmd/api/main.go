package main

import (
	"flag"
	"log/slog"
	"os"
)

type config struct {
	port int
	// in a real project this would also include environment
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
}
