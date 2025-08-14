package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	// in a real project, this would be wrapped in a goroutine to enable receiving
	// system interrupts and handling graceful shutdowns
	app.logger.Info("starting server", "addr", srv.Addr)
	return srv.ListenAndServe()
}
