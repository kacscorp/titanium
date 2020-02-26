package db

import "fmt"

// Configuration holds all the configuration items required for opening a new
// connection to a database.
type Configuration struct {
	Host               string `json:"host" yaml:"host"`
	MaxIdleConnections *int   `json:"max_idle_connections" yaml:"max_idle_connections"`
	MaxOpenConnections *int   `json:"max_open_connections" yaml:"max_open_connections"`
	Name               string `json:"name" yaml:"name"`
	Password           string `json:"password" yaml:"password"`
	Port               int    `json:"port" yaml:"port"`
	User               string `json:"user" yaml:"user"`
}

// Configurations is a group of Configurations, each with a unique label
type Configurations map[string]Configuration

func (cfg *Configuration) url() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)
}
