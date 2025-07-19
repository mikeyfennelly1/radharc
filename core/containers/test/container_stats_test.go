package containers

import (
	"github.com/mikeyfennelly1/radharc/core/containers"
	"testing"
)

func Test_GetContainerStats(t *testing.T) {
	t.Run("basic run func", func(t *testing.T) {
		// get a random container id
		runningContainers, _ := containers.ListRunningContainers()
		var containerID string
		for _, runningContainer := range runningContainers {
			containerID = runningContainer.ID
			break
		}
		containers.GetContainerStats(containerID)
	})
}
