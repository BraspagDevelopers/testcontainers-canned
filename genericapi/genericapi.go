package genericapi

import (
	"context"
	"fmt"

	canned "github.com/BraspagDevelopers/testcontainers-canned"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultPort             = "80/tcp"
	defaultLivenessEndpoint = "/"
)

// ContainerRequest a container request specification
type ContainerRequest struct {
	testcontainers.GenericContainerRequest
	Image            string
	LivenessEndpoint string
	Port             nat.Port
}

// Container represents a mock-server container
type Container struct {
	Container testcontainers.Container
	req       ContainerRequest
}

// WithNetworkAlias adds a network alias to the container request
func (req ContainerRequest) WithNetworkAlias(network, alias string) ContainerRequest {
	canned.AddNetworkAlias(&req.GenericContainerRequest, network, alias)
	return req
}

// CreateContainer creates and starts a container and assumes it contains an HTTP API
func CreateContainer(ctx context.Context, req ContainerRequest) (*Container, error) {
	if req.Port == "" {
		req.Port = defaultPort
	}
	if req.Image == "" {
		return nil, errors.New("an image name is required")
	}
	req.GenericContainerRequest.Image = req.Image
	if req.ExposedPorts == nil {
		req.ExposedPorts = []string{string(req.Port)}
	}
	if req.LivenessEndpoint == "" {
		req.LivenessEndpoint = defaultLivenessEndpoint
	}
	if req.WaitingFor == nil {
		req.WaitingFor = wait.ForHTTP(req.LivenessEndpoint).
			WithPort(req.Port)
	}

	provider, err := req.ProviderType.GetProvider()
	if err != nil {
		return nil, err
	}

	result := &Container{
		req: req,
	}

	req.Started = false
	if result.Container, err = provider.CreateContainer(ctx, req.ContainerRequest); err != nil {
		return result, errors.Wrap(err, "Error creating container")
	}

	if err = result.Container.Start(ctx); err != nil {
		return result, errors.Wrap(err, "Error starting container")
	}

	return result, nil
}

// URL builds an URL that can be used to interact with the container's HTTP API
func (c Container) URL(ctx context.Context) (string, error) {
	host, err := c.Container.Host(ctx)
	if err != nil {
		return "", errors.Wrap(err, "Error reading container host name")
	}
	port, err := c.Container.MappedPort(ctx, c.req.Port)
	if err != nil {
		return "", errors.Wrap(err, "Error reading container mapped port")
	}
	return fmt.Sprintf("http://%s:%s", host, port.Port()), nil
}

// URLForNetwork builds an URL that can be used to interact with the container's HTTP API inside the specified network
func (c Container) URLForNetwork(ctx context.Context, network string) (string, error) {

	alias, err := canned.GetAliasForNetwork(ctx, c.req.GenericContainerRequest, network)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://%s:%s", alias, c.req.Port.Port()), nil
}

// HostAndPort the host and port that can be used to interact with the container's HTTP API
func (c Container) HostAndPort(ctx context.Context) (string, nat.Port, error) {
	return canned.GetHostAndPort(ctx, c.Container, c.req.Port)
}
