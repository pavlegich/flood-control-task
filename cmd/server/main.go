// Package main contains actions for building and running the server.
package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/pavlegich/flood-control-task/internal/controllers"
	"github.com/pavlegich/flood-control-task/internal/domains/rwmanager"
	"github.com/pavlegich/flood-control-task/internal/infra/config"
	"github.com/pavlegich/flood-control-task/internal/infra/database"
	"github.com/pavlegich/flood-control-task/internal/infra/logger"
	"github.com/pavlegich/flood-control-task/internal/server"
	"github.com/pavlegich/flood-control-task/internal/utils"
	"go.uber.org/zap"
)

func main() {
	// Context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	// Manager for read and write
	rw := rwmanager.NewRWManager(ctx, os.Stdin, os.Stdout)

	// f, _ := os.Open("../../test/commands")
	// rw := rwmanager.NewRWManager(ctx, f, os.Stdout)

	// Greeting
	rw.Writeln(ctx, utils.Greet)
	// WaitGroup
	wg := &sync.WaitGroup{}

	// Logger
	err := logger.Init(ctx, "Info")
	if err != nil {
		logger.Log.Error("main: logger initialization failed", zap.Error(err))
	}
	defer logger.Log.Sync()

	// Configuration
	cfg := config.NewConfig(ctx)
	err = cfg.ParseFlags(ctx)
	if err != nil {
		logger.Log.Error("main: parse flags failed", zap.Error(err))
	}

	// Database
	db, err := database.Init(ctx, cfg.DSN)
	if err != nil {
		logger.Log.Error("main: database initialization failed", zap.Error(err))
	}
	defer db.Close()

	// Server
	ctrl := controllers.NewController(ctx, rw, cfg)
	server, err := server.NewServer(ctx, ctrl, rw, cfg)
	if err != nil {
		logger.Log.Error("main: create new server failed", zap.Error(err))
	}

	// Run Server
	wg.Add(1)
	go func() {
		err := server.Serve(ctx)
		if err != nil {
			logger.Log.Error("main: server serve error", zap.Error(err))
		}
		stop()
		wg.Done()
	}()

	// Server graceful shutdown
	<-ctx.Done()
	if ctx.Err() != nil {
		logger.Log.Info("shutting down gracefully...",
			zap.Error(ctx.Err()))

		connsClosed := make(chan struct{})
		go func() {
			wg.Wait()
			close(connsClosed)
		}()

		select {
		case <-connsClosed:
		case <-time.After(5 * time.Second):
			rw.Writeln(ctx, "\n"+utils.UnexpectedQuit)
			panic("shutdown timeout")
		}
	}

	rw.Writeln(ctx, utils.Quit)
}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
