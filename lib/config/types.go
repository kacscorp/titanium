package config

import (
	"time"

	"github.com/kacscorp/titanium/lib/config/db"
)

// Configuration contains the config information of the service and all of its
// components.
type Configuration struct {
	GracefulShutdownSeconds time.Duration     `json:"graceful_shutdown_seconds" yaml:"graceful_shutdown_seconds"`
	Port                    int               `json:"port" yaml:"port"`
	Databases               db.Configurations `json:"databases" yaml:"databases"`
}
