package flood

import "context"

type FloodController struct {
}

func (c *FloodController) Check(ctx context.Context, userID int64) (bool, error) {
	return true, nil
}
