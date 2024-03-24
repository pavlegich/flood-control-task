package flood

import (
	"context"

	"github.com/pavlegich/flood-control-task/internal/infra/config"
)

// FloodController contains storage for flood checking.
type FloodController struct {
	repo Repository
	cfg  *config.Config
}

// NewFloodController returns new FloodController object.
func NewFloodController(ctx context.Context, repo Repository, cfg *config.Config) *FloodController {
	return &FloodController{
		repo: repo,
		cfg:  cfg,
	}
}

// Check checks whether the limit of the maximum allowed number of requests has been reached
// by user according to the specified flood control rules.
func (c *FloodController) Check(ctx context.Context, userID int64) (bool, error) {

	return true, nil
}
