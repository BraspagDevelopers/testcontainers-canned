package genericapi

import (
	"testing"

	canned "github.com/BraspagDevelopers/testcontainers-canned"
)

func TestContainerMustBeContainer(t *testing.T) {
	var c *Container
	var x canned.Container
	x = c
	_ = x
}
