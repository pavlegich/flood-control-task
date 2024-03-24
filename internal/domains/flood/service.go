package flood

import (
	"context"
	"fmt"
	"sync"

	"github.com/pavlegich/flood-control-task/internal/infra/config"
)

// FloodController contains storage for calls checking.
type FloodController struct {
	repo Repository
	cfg  *config.Config
	mu   *sync.Mutex
}

// NewFloodController returns new FloodController object.
func NewFloodController(ctx context.Context, repo Repository, cfg *config.Config) *FloodController {
	return &FloodController{
		repo: repo,
		cfg:  cfg,
		mu:   &sync.Mutex{},
	}
}

// Check checks whether the limit of the maximum allowed number of calls has been reached
// by user according to the specified flood control rules.
func (c *FloodController) Check(ctx context.Context, userID int64) (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	count, err := c.repo.CreateCall(ctx, userID, c.cfg.Period)
	if err != nil {
		return false, fmt.Errorf("Check: couldn't get calls count %w", err)
	}
	if count > c.cfg.Calls {
		return false, nil
	}

	return true, nil
}
