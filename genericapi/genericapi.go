package genericapi

import (
	"context"
	"fmt"

	"github.com/BraspagDevelopers/canned-testcontainers"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultPort             = "80/tcp"
	defaultLivenessEndpoint = "/"
)

type ContainerRequest struct {
	testcontainers.GenericContainerRequest
	Image            string
	LivenessEndpoint string
	Port             nat.Port
}
type Container struct {
	Container testcontainers.Container
	req       ContainerRequest
}

func (req ContainerRequest) WithNetworkAlias(network, alias string) ContainerRequest {
	canned.AddNetworkAlias(&req.GenericContainerRequest, network, alias)
	return req
}

func (c Container) BaseURL(ctx context.Context) (string, error) {
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
	} else {
		req.WaitingFor = wait.ForAll(
			req.WaitingFor,
			wait.ForHTTP(req.LivenessEndpoint).
				WithPort(req.Port))
	}

	provider, err := req.ProviderType.GetProvider()
	if err != nil {
		return nil, err
	}

	result := &Container{
		req: req,
	}

	if result.Container, err = provider.CreateContainer(ctx, req.ContainerRequest); err != nil {
		return result, errors.Wrap(err, "Error creating container")
	}

	if err = result.Container.Start(ctx); err != nil {
		return result, errors.Wrap(err, "Error starting container")
	}

	return result, nil
}
