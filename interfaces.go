package canned

import (
	"context"
)

type Shutdownable interface {
	Shutdown(ctx context.Context) error
}

type Loggable interface {
	GetLogs(ctx context.Context) (string, error)
}
