package mockserver

import (
	"context"

	mockserverclient "github.com/BraspagDevelopers/mock-server-client"
	canned "github.com/BraspagDevelopers/testcontainers-canned"
	"github.com/BraspagDevelopers/testcontainers-canned/genericapi"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
)

const (
	image            = "mockserver/mockserver"
	exposedPort      = "1080/tcp"
	livenessEndpoint = "/_meta/alive"
)

// ContainerRequest a container request specification
type ContainerRequest genericapi.ContainerRequest

// Container represents a mock-server container
type Container genericapi.Container

// WithNetworkAlias adds a network alias to the container request
func (req ContainerRequest) WithNetworkAlias(network, alias string) ContainerRequest {
	canned.AddNetworkAlias(&req.GenericContainerRequest, network, alias)
	return req
}

// CreateContainer creates and starts a mock-server container
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

// URL builds an URL that can be used to interact with the container's HTTP API
func (c Container) URL(ctx context.Context) (string, error) {
	return genericapi.Container(c).URL(ctx)
}

// URLForNetwork builds an URL that can be used to interact with the container's HTTP API inside the specified network
func (c Container) URLForNetwork(ctx context.Context, network string) (string, error) {
	return genericapi.Container(c).URLForNetwork(ctx, network)
}

// HostAndPort retrieves the external host and port of the container
func (c Container) HostAndPort(ctx context.Context) (string, nat.Port, error) {
	return genericapi.Container(c).HostAndPort(ctx)
}

// Client creates and returns a new mock-server client
func (c Container) Client(ctx context.Context) (*mockserverclient.MockServerClient, error) {
	url, err := c.URL(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not get the mock server URL")
	}
	client := mockserverclient.NewClientURL(url)
	return &client, nil
}
