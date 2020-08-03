package canned

import (
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

// // GetLogs retrieves all container logs
// func GetLogs(ctx context.Context, c testcontainers.Container) (string, error) {
// 	var sb strings.Builder
// 	if reader, err := c.Logs(ctx); err == nil {
// 		buf := new(bytes.Buffer)
// 		if _, err := buf.ReadFrom(reader); err == nil {
// 			sb.Write(buf.Bytes())
// 			sb.WriteRune('\n')
// 		} else {
// 			return "", errors.Wrap(err, "cannot read logs")
// 		}
// 	} else {
// 		return "", errors.Wrap(err, "cannot read logs")
// 	}
// 	return sb.String(), nil
// }

type _TestingLogger struct {
	Logger LoggerWithLog
}

type LoggerWithLog interface {
	Log(v ...interface{})
}

func NewTestingLogger(logger LoggerWithLog) *testcontainers.LogConsumer {
	var lc testcontainers.LogConsumer
	lc = _TestingLogger{
		Logger: logger,
	}
	return &lc
}
func (c _TestingLogger) Accept(e testcontainers.Log) {
	c.Logger.Log(fmt.Sprintf("%s: %s", e.LogType, string(e.Content)))
}
