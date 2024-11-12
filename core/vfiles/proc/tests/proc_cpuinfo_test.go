package vfiles

import (
	vfiles "github.com/mikeyfennelly1/radharc/core/vfiles/proc"
	//"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cpuinfo(t *testing.T) {
	t.Run("Run the file", func(t *testing.T) {
		cpuInfo := vfiles.GetCPUInfo()

		println(cpuInfo)
	})
}
