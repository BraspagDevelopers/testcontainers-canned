package canned

import (
	"bytes"
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
)

// GetLogs retrieves all logs from the container
func GetLogs(ctx context.Context, c testcontainers.Container) (string, error) {
	var sb strings.Builder
	if reader, err := c.Logs(ctx); err == nil {
		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(reader); err == nil {
			sb.Write(buf.Bytes())
			sb.WriteRune('\n')
		} else {
			return "", errors.Wrap(err, "cannot read logs")
		}
	} else {
		return "", errors.Wrap(err, "cannot read logs")
	}
	return sb.String(), nil
}
