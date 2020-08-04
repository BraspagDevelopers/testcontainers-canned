package mockserver

import (
	"context"

	canned "github.com/BraspagDevelopers/testcontainers-canned"
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
