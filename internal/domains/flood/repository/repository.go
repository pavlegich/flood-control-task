// Package repository contains repository object
// and methods for interaction with storage.
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Repository contains storage objects.
type Repository struct {
	db *sql.DB
}

// NewFloodRepository returns new repository object.
func NewFloodRepository(ctx context.Context, db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateCall gets and returns user's calls count in the last requested seconds.
func (r *Repository) CreateCall(ctx context.Context, userID int64, limit time.Duration) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return -1, fmt.Errorf("CreateCall: begin transaction failed %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `INSERT INTO flood (user_id) VALUES ($1)`, userID)
	if err != nil {
		return -1, fmt.Errorf("CreateCall: insert new call failed %w", err)
	}

	row := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM flood WHERE user_id = $1 AND 
		created_at > $2`, userID, time.Now().Add(-limit))

	var callsCount int
	err = row.Scan(&callsCount)
	if err != nil {
		return -1, fmt.Errorf("CreateCall: scan calls count for user failed %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return -1, fmt.Errorf("CreateCall: commit transaction failed %w", err)
	}

	return callsCount, nil
}
