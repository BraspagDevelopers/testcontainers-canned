package networks

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
)

// Network represents a network
type Network struct {
	testcontainers.Network
	Name string
}

// CreateNetwork creates a new network with a random name
func CreateNetwork(ctx context.Context) (*Network, error) {
	var network Network
	network.Name = fmt.Sprintf("net_%s", uuid.New())
	n, err := testcontainers.GenericNetwork(ctx, testcontainers.GenericNetworkRequest{
		NetworkRequest: testcontainers.NetworkRequest{
			Name:           network.Name,
			CheckDuplicate: true,
		},
	})
	network.Network = n
	return &network, err
}
