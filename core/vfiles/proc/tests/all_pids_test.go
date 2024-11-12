package vfiles

import (
	"fmt"
	vfiles "github.com/mikeyfennelly1/radharc/core/vfiles/proc"
	"testing"
)

func Test_GetAllRunningPids(t *testing.T) {
	t.Run("run func", func(t *testing.T) {
		runningPids := vfiles.GetAllRunningPids()
		for _, pid := range runningPids {
			fmt.Printf("%d\n", pid)
		}
	})
}
