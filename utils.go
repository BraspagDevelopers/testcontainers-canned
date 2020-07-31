package canned

import "github.com/testcontainers/testcontainers-go"

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
