package sqlserver

import (
	"context"

	canned "github.com/BraspagDevelopers/testcontainers-canned"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
)

// Shutdown terminates the container
func (c *Container) Shutdown(ctx context.Context) error {
	if c != nil && c.Container != nil {
		return c.Container.Terminate(ctx)
	}
	return nil
}

// GetLogs retrieves all logs from the container
func (c *Container) GetLogs(ctx context.Context) (string, error) {
	if c != nil && c.Container != nil {
		return canned.GetLogs(ctx, c.Container)
	}
	return "", nil
}

// HostAndPort retrieves the external host and port of the container
func (c *Container) HostAndPort(ctx context.Context) (string, nat.Port, error) {
	if c != nil && c.Container != nil {
		return canned.GetHostAndPort(ctx, c.Container, exposedPort)
	}
	return "", "", errors.New("could not read host and port from a nil pointer")
}
