package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	defaultDriver                  = "postgres"
	msgFmtFailedToOpenDBConnection = "failed to open DB connection, url='%s'"
	msgFmtPingTestFailed           = "ping test failed, url='%s'"
	pkgName                        = "db"
)

// Connect opens a sql connection and returns the pool
func Connect(cfg Configuration) (*gorm.DB, error) {
	conn, err := gorm.Open(defaultDriver, cfg.url())
	if err != nil {
		return nil, errors.New(msgFmtFailedToOpenDBConnection)
	}

	fmt.Println("Successfully connected!")
	return conn, nil
}
