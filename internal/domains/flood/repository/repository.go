// Package repository contains repository object
// and methods for interaction with storage.
package repository

import (
	"context"
	"database/sql"
	"time"
	// "github.com/jackc/pgerrcode"
	// "github.com/pavlegich/flood-control-task/internal/domains/flood"
)

// Repository contains storage objects.
type Repository struct {
	db *sql.DB
}

// NewDataRepository returns new repository object.
func NewDataRepository(ctx context.Context, db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// GetReqCountByUserID gets and returns user's requests count in the last requested seconds
func (r *Repository) GetReqCountByUserID(ctx context.Context, userID int64, limit time.Time) (int, error) {
	return 0, nil
}
