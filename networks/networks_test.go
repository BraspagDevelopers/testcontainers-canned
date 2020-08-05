package networks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAndShutdown(t *testing.T) {
	ctx := context.Background()

	n, err := CreateNetwork(ctx)
	require.NoError(t, err, "create error")

	err = n.Shutdown(ctx)
	require.NoError(t, err, "shutdown error")
}
