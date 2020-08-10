package canned

import (
	"context"

	"github.com/docker/go-connections/nat"
)

// Shutdownable represents something that can be shut down
type Shutdownable interface {
	Shutdown(ctx context.Context) error
}

// Container represents a container
type Container interface {
	Shutdownable
	GetLogs(context.Context) (string, error)
	HostAndPort(context.Context) (string, nat.Port, error)
}
