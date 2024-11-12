package containers

import (
	"github.com/mikeyfennelly1/radharc/core/containers"
	"testing"
)

func Test_ListContainers(t *testing.T) {
	t.Run("test list containers", func(t *testing.T) {
		containers.ListRunningContainers()
	})
}

func Test_ListImages(t *testing.T) {
	t.Run("test list containers", func(t *testing.T) {
		containers.ListAllImages()
	})
}
