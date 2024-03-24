// Package utils contains additional methods for server.
package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/pavlegich/flood-control-task/internal/domains/rwmanager"
	errs "github.com/pavlegich/flood-control-task/internal/errors"
)

const (
	// Constants contain phrases for server output.
	Greet          = "Welcome!"
	Quit           = "quit"
	Success        = "success"
	Exit           = "exit"
	Close          = "close"
	UnexpectedQuit = "unexpected quit"
)

// GetKnownErr checks the error and returns it, if it is known.
func GetKnownErr(err error) error {
	if errors.Is(err, errs.ErrLimitExceeded) {
		return errs.ErrLimitExceeded
	}
	if errors.Is(err, errs.ErrUnknownCommand) {
		return errs.ErrUnknownCommand
	}
	return nil
}

// DoWithRetryIfEmpty tries to implement function three times, if the input is empty.
func DoWithRetryIfEmpty(ctx context.Context, rw rwmanager.RWService, f func(ctx context.Context) error) error {
	var err error
	for i := 0; i < 3; i++ {
		err = f(ctx)
		if ctx.Err() != nil {
			return fmt.Errorf("DoWithRetryIfEmpty: context error %w", ctx.Err())
		}
		if !errors.Is(err, errs.ErrEmptyInput) {
			return err
		}
		rw.Error(ctx, errs.ErrEmptyInput)
	}
	return err
}
