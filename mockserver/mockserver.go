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
