package mockserver

import (
	"context"

	"github.com/BraspagDevelopers/canned-testcontainers"
	"github.com/BraspagDevelopers/canned-testcontainers/genericapi"
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
