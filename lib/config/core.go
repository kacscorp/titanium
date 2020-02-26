package config

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kacscorp/titanium/lib/config/db"
	"github.com/kataras/muxie"
)

const (
	msgFailedToCreateMultiplexer = "failed to create a httpmux.IMultiplexer"
	msgFailedToCreateRouter      = "failed to create a router"
	msgFmtFailedToConnectToDB    = "failed to connect to '%s' database"
	pkgName                      = "core"
)

// Run starts the service. It receives a configuration object which it uses for
// initializing all the dependencies.
//
// In case of error, it is returned.
func Run(config *Configuration) error {

	// Dependency: Multiplexer
	mux := muxie.NewMux()

	// Dependency: Titanium Database connection
	titaniumDB, err := connectToDB(config, "titanium_database")
	if err != nil {
		return err
	}

	logger := log.Logger{}
	// Server
	server, err := newServer(titaniumDB, mux, &logger)
	if err != nil {
		return errors.New(msgFailedToCreateRouter)
	}
	// gService.ShutdownAtExit(server)

	server.start()

	return nil
}

func connectToDB(config *Configuration, dbName string) (*gorm.DB, error) {
	conn, err := db.Connect(config.Databases[dbName])
	if err != nil {
		return nil, errors.New(msgFmtFailedToConnectToDB)
	}

	return conn, nil
}
