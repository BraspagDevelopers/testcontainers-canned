package genericapi

import (
	"bytes"
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWhenDoesNotProvideImage_ShouldReturnError(t *testing.T) {
	ctx := context.Background()

	c, err := New(ctx, ContainerRequest{})
	assert.EqualError(t, err, "an image name is required")
	assert.Nil(t, c)
}

func TestEcho(t *testing.T) {
	ctx := context.Background()

	c, err := New(ctx, ContainerRequest{
		Image: "jmalloc/echo-server",
		Port:  "8080/tcp",
	})
	require.NoError(t, err)

	url, err := c.BaseURL(ctx)
	require.NoError(t, err)

	resp, err := http.Post(url, "text/plain", strings.NewReader("Hello World!"))
	require.NoError(t, err)
	assert.NotNil(t, resp)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()
	assert.Contains(t, body, "Hello World!")
}
