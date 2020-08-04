package sqlserver

import (
	"testing"

	canned "github.com/BraspagDevelopers/testcontainers-canned"
)

func TestMustBeShutdownable(t *testing.T) {
	var c *Container
	var x canned.Shutdownable
	x = c
	_ = x
}

func TestMustBeLoggable(t *testing.T) {
	var c *Container
	var x canned.Loggable
	x = c
	_ = x
}
