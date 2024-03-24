// Package flood contains object and methods for
// checking the user flood.
package flood

import (
	"context"
	"time"
)

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}

// Repository describes methods related with FloodControl object
// for interaction with database.
type Repository interface {
	CreateCall(ctx context.Context, userID int64, limit time.Duration) (int, error)
}
