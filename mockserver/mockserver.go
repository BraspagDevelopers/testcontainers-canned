package mockserver

import (
	"context"

	canned "github.com/BraspagDevelopers/testcontainers-canned"
	"github.com/BraspagDevelopers/testcontainers-canned/genericapi"
)

const (
	image            = "mockserver/mockserver"
	exposedPort      = "1080/tcp"
	livenessEndpoint = "/_meta/alive"
)

type ContainerRequest genericapi.ContainerRequest
type Container genericapi.Container

func (req ContainerRequest) WithNetworkAlias(network, alias string) ContainerRequest {
	canned.AddNetworkAlias(&req.GenericContainerRequest, network, alias)
	return req
}

func (c Container) URL(ctx context.Context) (string, error) {
	return genericapi.Container(c).URL(ctx)
}

func (c Container) URLForNetwork(ctx context.Context, network string) (string, error) {
	return genericapi.Container(c).URLForNetwork(ctx, network)
}

func CreateContainer(ctx context.Context, req ContainerRequest) (*Container, error) {
	if req.Port == "" {
		req.Port = exposedPort
	}
	if req.Image == "" {
		req.Image = image
	}
	if req.LivenessEndpoint == "" {
		req.LivenessEndpoint = livenessEndpoint
	}
	apic, err := genericapi.CreateContainer(ctx, genericapi.ContainerRequest(req))
	c := Container(*apic)
	return &c, err
}

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
