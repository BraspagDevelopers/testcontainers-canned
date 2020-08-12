package mockserver

import (
	"context"
	"errors"

	"github.com/BraspagDevelopers/testcontainers-canned/genericapi"
	"github.com/docker/go-connections/nat"
)

// Shutdown terminates the container
func (c Container) Shutdown(ctx context.Context) error {
	if c.Container != nil {
		return genericapi.Container(c).Shutdown(ctx)
	}
	return nil
}

// GetLogs retrieves all logs from the container
func (c Container) GetLogs(ctx context.Context) (string, error) {
	if c.Container != nil {
		return genericapi.Container(c).GetLogs(ctx)
	}
	return "", nil
}

// HostAndPort retrieves the external host and port of the container
func (c Container) HostAndPort(ctx context.Context) (string, nat.Port, error) {
	if c.Container != nil {
		return genericapi.Container(c).HostAndPort(ctx)
	}
	return "", "", errors.New("could not read host and port from a nil pointer")
}
