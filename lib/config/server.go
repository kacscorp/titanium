package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kacscorp/titanium/lib/config/handlers"
	"github.com/kacscorp/titanium/lib/sources/employees/v1"
	"github.com/kataras/muxie"
	"github.com/zenazn/goji/graceful"
)

const (
	msgFmtFailedToRegisterRoute = "failed to register /%s route"
)

// router is basically a wrapper around a web.Mux
type server struct {
	titaniumDB *gorm.DB
	mux        *muxie.Mux
	logger     *log.Logger
}

// newServer builds and returns a pointer to a router
func newServer(
	titaniumDB *gorm.DB,
	mux *muxie.Mux,
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
	//Titanium Employees endpoints
	// Employee GET endpoint
	if handler, err := handlers.NewUsingSourceHandler(context, employees.GetHandler); err != nil {
		fmt.Errorf("Employee handler error")
	} else {
		sv.mux.Handle("/employees", handler)
	}
}

// start starts the router by telling its httpmux.IMultiplexer to Serve()
func (sv *server) start() {
	sv.defineRoutes()
	graceful.ListenAndServe(":8000", sv.mux)
}
