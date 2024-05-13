package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cristianortiz/htmxTempl-Go/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

// init() is called before main, ideal to load env vars before anything else
func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file, using default values.")
	}
}

func main() {
	router := chi.NewMux()
	//Make() is HTTPHandler wrapper
	router.Get("/foo", handlers.Make(handlers.HandleFoo))
	//Bind address for the server
	port := os.Getenv("SERVER_PORT")

	// create a new server
	s := http.Server{
		Addr:    port,   // configure port
		Handler: router, // set the default handler
		//ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server using goroutine
	go func() {
		slog.Info("Starting server", "listenAddr", port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Error starting server: ", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	// Block until a signal is received.
	sig := <-quit
	slog.Info("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	if err := s.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed: ", err)
	}
	slog.Info("Server shutdown gracefully")

	//explicit cancel() in context.WithTimeout and defer it later to avoid warning about cancel function not discarded
	defer cancel()
}
