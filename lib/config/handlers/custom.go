package handlers

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// AppContext contains our local context; our database pool, session store, template
// registry and anything else our handlers need to access. We'll create an instance of it
// in our main() function and then explicitly pass a reference to it for our handlers to access.
type AppContext struct {
	DB *gorm.DB
	// ... and the rest of our globals.
}

// We've turned our original AppHandler into a struct with two fields:
// - A function type similar to our original handler type (but that now takes an *AppContext)
// - An embedded field of type *AppContext
type AppHandler struct {
	Context *AppContext
	Handler func(*AppContext, http.ResponseWriter, *http.Request) (int, error)
}

// Our ServeHTTP method is mostly the same, and also has the ability to
// access our *AppContext's fields (templates, loggers, etc.) as well.
func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Updated to pass ah.AppContext as a parameter to our handler type.
	status, err := ah.Handler(ah.Context, w, r)
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
