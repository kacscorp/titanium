package config

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"

	"github.com/kacscorp/titanium/lib/config/handlers"
	"github.com/kacscorp/titanium/lib/sources/employees/v1"
)

const (
	msgFmtFailedToRegisterRoute = "failed to register /%s route"
)

// router is basically a wrapper around a web.Mux
type server struct {
	titaniumDB *gorm.DB
	mux        *web.Mux
	logger     *log.Logger
}

// newServer builds and returns a pointer to a router
func newServer(
	titaniumDB *gorm.DB,
	mux *web.Mux,
	logger *log.Logger,
) (*server, error) {
	return &server{
		titaniumDB: titaniumDB,
		mux:        mux,
		logger:     logger,
	}, nil
}

func (sv *server) defineRoutes() {

	context := &handlers.AppContext{DB: sv.titaniumDB}
	//Titanium Employees GET endpoint
	sv.mux.Get("/employees", handlers.AppHandler{Context: context, Handler: employees.GetHandler})
}

// start starts the router by telling its httpmux.IMultiplexer to Serve()
func (sv *server) start() {
	sv.defineRoutes()
	graceful.ListenAndServe(":8000", sv.mux)
}
