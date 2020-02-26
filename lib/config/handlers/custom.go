package handlers

import (
	"log"
	"net/http"
)

// custom is an object that satisfies http.Handler. It takes a function matching
// the signature that can return errors and a logger which it uses for logging
// when the handler fails.
type custom struct {
	handlerFn errorReturningHandlerFunc
}

// errorReturningHandlerFunc is a function that has a similar signature to an
// http.Handler but can return an error.
type errorReturningHandlerFunc func(http.ResponseWriter, *http.Request) (int, error)

// usingContextHandlerFunc is a function that has a similar signature to an
// http.Handler but can receive a Context and returns an error.
type usingContextHandlerFunc func(*AppContext, http.ResponseWriter, *http.Request) (int, error)

// New receives a base handler function that returns errors and
// also receives a logger. With that, it builds a custom handler and returns a
// pointer to it.
//
// The idea of this is to be able to write a new type of handler with a
// signature that allows them to send errors upstream. All logging occurs at a
// higher level so the logger doesn't need to be injected everywhere. With this,
// it is very important to wrap the errors (see errutil package), so they
// contain enough description about what happened.
func New(handlerFn errorReturningHandlerFunc) (IHTTPHandler, error) {
	return &custom{
		handlerFn: handlerFn,
	}, nil
}

// NewUsingSourceHandler receives a *sql.DB, a handler function that receives a
// *sql.DB and returns errors and a also receives a logger. With that, it builds
// a custom handler (the one that can return errors) and returns a pointer to
// it. The function passed to this custom handler is one that makes use of the
// *sql.DB received.
func NewUsingSourceHandler(context *AppContext, handlerFn usingContextHandlerFunc) (IHTTPHandler, error) {
	return New(func(w http.ResponseWriter, r *http.Request) (int, error) {
		return handlerFn(context, w, r)
	})
}

// Our ServeHTTP method is mostly the same, and also has the ability to
// access our *AppContext's fields (templates, loggers, etc.) as well.
func (cst custom) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Updated to pass AppContext as a parameter to our handler type.
	status, err := cst.handlerFn(w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
			// And if we wanted a friendlier error page:
			// err := ah.renderTemplate(w, "http_404.tmpl", nil)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}
