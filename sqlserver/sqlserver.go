package sqlserver

import (
	"context"
	"fmt"
	"time"

	"github.com/BraspagDevelopers/canned-testcontainers"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	username    = "sa"
	image       = "mcr.microsoft.com/mssql/server"
	exposedPort = "1433/tcp"
)

type ContainerRequest struct {
	testcontainers.GenericContainerRequest
	Username string
	Password string
	Image    string
}

type Container struct {
	Container testcontainers.Container
	req       ContainerRequest
}

func (req ContainerRequest) WithNetworkAlias(network, alias string) ContainerRequest {
	canned.AddNetworkAlias(&req.GenericContainerRequest, network, alias)
	return req
}

func (c *Container) GoConnectionString(ctx context.Context) (string, error) {
	host, port, err := canned.GetHostAndPort(ctx, c.Container, exposedPort)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s", c.req.Username, c.req.Password, host, port.Port()), nil
}

func (c *Container) DotNetConnectionStringForNetwork(ctx context.Context, network string) (string, error) {
	alias, err := canned.GetAliasForNetwork(ctx, c.req.GenericContainerRequest, network)
	if err != nil {

	}
	return fmt.Sprintf("Data Source=%s,%s; User ID=%s; Password=%s; Pooling=False;", alias, exposedPort, c.req.Username, c.req.Password), nil
}

func New(ctx context.Context, req ContainerRequest) (*Container, error) {
	if req.Image == "" {
		req.Image = image
	}
	req.GenericContainerRequest.Image = req.Image
	if req.ExposedPorts == nil {
		req.ExposedPorts = []string{exposedPort}
	}

	if req.Username == "" {
		req.Username = username
	}
	if req.Password == "" {
		return nil, errors.New("A password must be provided")
	}
	if req.WaitingFor == nil {
		req.WaitingFor = wait.ForSQL(exposedPort, "sqlserver", func(port nat.Port) string {
			return fmt.Sprintf("sqlserver://%s:%s@localhost:%s", req.Username, req.Password, port.Port())
		}).Timeout(time.Minute)
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
