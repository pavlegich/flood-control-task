// Package server contains Server object and its methods.
package server

import (
	"context"
	"fmt"

	"github.com/pavlegich/flood-control-task/internal/controllers"
	"github.com/pavlegich/flood-control-task/internal/domains/rwmanager"
	"github.com/pavlegich/flood-control-task/internal/infra/config"
	"github.com/pavlegich/flood-control-task/internal/utils"
)

// Server contains client attributes.
type Server struct {
	rw         rwmanager.RWService
	controller *controllers.Controller
	config     *config.Config
}

// NewServer initializes controller and router, returns new server object.
func NewServer(ctx context.Context, ctrl *controllers.Controller, rw rwmanager.RWService, cfg *config.Config) (*Server, error) {
	return &Server{
		controller: ctrl,
		rw:         rw,
		config:     cfg,
	}, nil
}

// Serve starts listening and catching the commands from standart input.
func (s *Server) Serve(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err := utils.DoWithRetryIfEmpty(ctx, s.rw, s.controller.HandleCommand)
			if err != nil {
				got := utils.GetKnownErr(err)
				if got == nil {
					return fmt.Errorf("Serve: handle command failed %w", err)
				}
				s.rw.Error(ctx, got)
			}
		}
	}
}
