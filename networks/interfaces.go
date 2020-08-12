package networks

import "context"

// Shutdown shuts the network down
func (n Network) Shutdown(ctx context.Context) error {
	if n.Network != nil {
		return n.Network.Remove(ctx)
	}
	return nil
}
