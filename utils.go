package canned

import (
	"context"
	"strings"

	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
)

func AddNetworkAlias(req *testcontainers.GenericContainerRequest, network, alias string) {
	if req.Networks == nil {
		req.Networks = make([]string, 0)
	}
	req.Networks = append(req.Networks, network)

	if req.NetworkAliases == nil {
		req.NetworkAliases = make(map[string][]string)
	}
	if _, ok := req.NetworkAliases[network]; !ok {
		req.NetworkAliases[network] = make([]string, 0)
	}
	req.NetworkAliases[network] = append(req.NetworkAliases[network], alias)
}

func GetHostAndPort(ctx context.Context, c testcontainers.Container, exposedPort nat.Port) (host string, port nat.Port, err error) {
	if host, err = c.Host(ctx); err != nil {
		err = errors.Wrap(err, "Error reading container host name")
		return
	}
	if port, err = c.MappedPort(ctx, exposedPort); err != nil {
		err = errors.Wrap(err, "Error reading container mapped port")
		return
	}
	return
}

func GetAliasForNetwork(ctx context.Context, req testcontainers.GenericContainerRequest, network string) (string, error) {
	hasNetwork := false
	for _, n := range req.Networks {
		if strings.EqualFold(n, network) {
			hasNetwork = true
			break
		}
	}
	if !hasNetwork {
		return "", errors.New("the container is not in the specified network")
	}

	var aliases []string
	var ok bool
	if aliases, ok = req.NetworkAliases[network]; !ok {
		return "", errors.New("the container is does not have an alias in the specified network")
	}
	if len(aliases) == 0 {
		return "", errors.New("the container is does not have an alias in the specified network")
	}
	return aliases[0], nil
}
