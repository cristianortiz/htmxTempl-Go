package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

// function type that takes two args, and returns an error type
// w used to write the response headers and body to the client
// r, contains info about the incoming request, ex method, URL, headers, body
// this is usfeul as a type alias for this function type, creating a more
// descriptive and reusable waay to represent handlers, in this case capable to handler http requests
// and return error in case of, HTTPHandler type allows the handlers creation with
// different functionalites whith a commmon interface
type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

// Make takes an HTTPHandler as input an returns an http.HandlerFunc, wich is
// another built-ind function type in the net/http pkg, it represents a function that
// implements the http.Handler INTERFACE.  So Make, is a wrapper around  HTTPHandler that creates and return
// a new function(anonymous), this new function acts like as any other HTTP handler ()
// the wrapper function centralizes error handling ensuring consisten loggin or the error
// managment across all handlers
func Make(h HTTPHandler) http.HandlerFunc {
	//the inner function takes h parameter (HTTPHandler received as input by Make())
	//and, if HTTPHandler returns an error, is registered with slog pkg
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
		}

	}
}

// helper func to simplify templ component rendering at integrating it with a http server,
// again this make http handlers more concise and easy to understand
func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	//calling Render() from templ.Component and pass it the request context
	//and http.ResponseWritter object,(w), c.Render renderize the c and writes de outpu in w
	return c.Render(r.Context(), w)
}
