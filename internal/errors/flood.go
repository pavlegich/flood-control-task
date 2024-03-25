package errors

import "errors"

var (
	ErrExit           = errors.New("exit requested")
	ErrUnknownCommand = errors.New("unknown command")
	ErrEmptyInput     = errors.New("input is empty")
	ErrLimitExceeded  = errors.New("calls limit exceeded, try again later")
)
