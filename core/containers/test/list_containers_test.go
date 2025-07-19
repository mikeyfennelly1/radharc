package containers

import (
	"fmt"
	"github.com/mikeyfennelly1/radharc/core/containers"
	"testing"
)

func Test_ListContainers(t *testing.T) {
	t.Run("test list containers", func(t *testing.T) {
		containers, _ := containers.ListRunningContainers()
		fmt.Printf("shshs")
		fmt.Println(containers)
	})
}

func Test_ListImages(t *testing.T) {
	t.Run("test list containers", func(t *testing.T) {
		containers.GetImages()
	})
}
