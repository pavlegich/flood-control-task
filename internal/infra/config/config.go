// Package config contains server configuration
// objects and methods
package config

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

// Config contains values of server flags and environments.
type Config struct {
	DSN    string        `env:"DATABASE_DSN" json:"database_dsn"`
	Period time.Duration `env:"PERIOD" json:"period"`
	Calls  int           `env:"CALLS" json:"calls"`
	UserID int64         `end:"USER_ID" json:"user_id"`
}

// NewConfig returns new server config.
func NewConfig(ctx context.Context) *Config {
	return &Config{}
}

// ParseFlags handles and processes flags and environments values
// when launching the server.
func (cfg *Config) ParseFlags(ctx context.Context) error {
	flag.StringVar(&cfg.DSN, "d", "postgresql://localhost:5432/flood", "URI (DSN) to database")
	flag.DurationVar(&cfg.Period, "p", time.Duration(60)*time.Second, "Time period for calls limit")
	flag.IntVar(&cfg.Calls, "c", 20, "Calls limit during the time period")
	flag.Int64Var(&cfg.UserID, "id", 1, "User ID for implementing the flood controller methods")

	flag.Parse()

	err := env.Parse(cfg)
	if err != nil {
		return fmt.Errorf("ParseFlags: wrong environment values %w", err)
	}

	return nil
}
