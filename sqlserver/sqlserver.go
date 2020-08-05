package sqlserver

import (
	"context"
	"fmt"
	"log"
	"time"

	canned "github.com/BraspagDevelopers/testcontainers-canned"
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

// ContainerRequest a container request specification
type ContainerRequest struct {
	testcontainers.GenericContainerRequest
	Username string
	Password string
	Image    string
	Logger   *testcontainers.LogConsumer
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

// CreateContainer creates a SQL Server for Linux container
func CreateContainer(ctx context.Context, req ContainerRequest) (*Container, error) {
	if req.Env == nil {
		req.Env = make(map[string]string)
	}
	req.Env["ACCEPT_EULA"] = "Y"
	if req.Image == "" {
		req.Image = image
		req.Env["MSSQL_PID"] = "Express"
	}
	req.GenericContainerRequest.Image = req.Image
	if req.ExposedPorts == nil {
		req.ExposedPorts = []string{exposedPort}
	}

	if req.Username == "" {
		req.Username = username
	}
	if req.Password == "" {
		return nil, errors.New("a password must be provided")
	} else {
		req.Env["SA_PASSWORD"] = req.Password
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

	req.Started = false
	log.Println("creating sqlserver")
	if result.Container, err = provider.CreateContainer(ctx, req.ContainerRequest); err != nil {
		return result, errors.Wrap(err, "could not create container")
	}

	if req.Logger != nil {
		log.Println("starting logger for sqlserver")
		if err = result.Container.StartLogProducer(ctx); err != nil {
			return result, errors.Wrap(err, "could not start log producer")
		}
		result.Container.FollowOutput(*req.Logger)
	}

	log.Println("starting sqlserver")
	if err = result.Container.Start(ctx); err != nil {
		return result, errors.Wrap(err, "could not start container")
	}
	return result, nil
}

// GoConnectionString returns a connection string suitable for usage in Go
func (c *Container) GoConnectionString(ctx context.Context) (string, error) {
	host, port, err := canned.GetHostAndPort(ctx, c.Container, exposedPort)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s", c.req.Username, c.req.Password, host, port.Port()), nil
}

// GoConnectionString returns a connection string suitable for usage in .NET
func (c *Container) DotNetConnectionStringForNetwork(ctx context.Context, network string) (string, error) {
	alias, err := canned.GetAliasForNetwork(ctx, c.req.GenericContainerRequest, network)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Data Source=%s,%s; User ID=%s; Password=%s; Pooling=False;", alias, exposedPort, c.req.Username, c.req.Password), nil
}
