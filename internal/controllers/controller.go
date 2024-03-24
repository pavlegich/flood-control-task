// Package controllers contains server controller object and
// methods for server work.
package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/pavlegich/flood-control-task/internal/domains/flood"
	repo "github.com/pavlegich/flood-control-task/internal/domains/flood/repository"
	"github.com/pavlegich/flood-control-task/internal/domains/rwmanager"
	errs "github.com/pavlegich/flood-control-task/internal/errors"
	"github.com/pavlegich/flood-control-task/internal/infra/config"
)

// Controller contains configuration for building the app.
type Controller struct {
	rw    rwmanager.RWService
	cfg   *config.Config
	flood flood.FloodControl
}

// NewController creates and returns new server controller.
func NewController(ctx context.Context, rw rwmanager.RWService, db *sql.DB, cfg *config.Config) *Controller {
	flood := flood.NewFloodController(ctx, repo.NewDataRepository(ctx, db), cfg)

	return &Controller{
		rw:    rw,
		cfg:   cfg,
		flood: flood,
	}
}

// HandleCommand handles commands from the input and does it, if the requested action is correct.
func (c *Controller) HandleCommand(ctx context.Context) error {
	c.rw.Write(ctx, "Type the command 'check', or exit: ")
	act, err := c.rw.Read(ctx)
	if err != nil {
		return fmt.Errorf("HandleCommand: read command failed %w", err)
	}

	act = strings.ToLower(act)
	switch act {
	case "check":
		ok, err := c.flood.Check(ctx, 1)
		if err != nil {
			return fmt.Errorf("HandleCommand: %w", err)
		}
		if !ok {
			return fmt.Errorf("HandleCommand: %w", errs.ErrLimitExceeded)
		}
	case "exit":
		return fmt.Errorf("HandleCommand: %w", errs.ErrExit)
	default:
		return fmt.Errorf("HandleCommand: %w", errs.ErrUnknownCommand)
	}

	return nil
}
