package networks

import (
	"testing"

	canned "github.com/BraspagDevelopers/testcontainers-canned"
)

func TestMustBeShutdownable(t *testing.T) {
	var n *Network
	var x canned.Shutdownable
	x = n
	_ = x
}
