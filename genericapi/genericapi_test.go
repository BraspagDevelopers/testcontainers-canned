package genericapi

import (
	"bytes"
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWhenDoesNotProvideImage_ShouldReturnError(t *testing.T) {
	ctx := context.Background()

	c, err := CreateContainer(ctx, ContainerRequest{})
	assert.EqualError(t, err, "an image name is required")
	assert.Nil(t, c)
}

func TestEcho(t *testing.T) {
	ctx := context.Background()

	c, err := CreateContainer(ctx, ContainerRequest{
		Image: "jmalloc/echo-server",
		Port:  "8080/tcp",
	})
	require.NoError(t, err)

	url, err := c.URL(ctx)
	require.NoError(t, err)

	resp, err := http.Post(url, "text/plain", strings.NewReader("Hello World!"))
	require.NoError(t, err)
	assert.NotNil(t, resp)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()
	assert.Contains(t, body, "Hello World!")

	err = c.Shutdown(ctx)
	require.NoError(t, err)
}

func TestWhenPortIsWrongWaitShouldThrowError(t *testing.T) {
	ctx := context.Background()

	req := ContainerRequest{
		Image: "braspagbrs.azurecr.io/canais/mocks/api/bpauth",
		Port:  "7000/tcp",
	}.WithNetworkAlias("my_network", "bpauth_api")
	req.WaitingFor = wait.ForHTTP("/live").WithStartupTimeout(3 * time.Second)

	_, err := CreateContainer(ctx, req)
	require.Error(t, err)
}
