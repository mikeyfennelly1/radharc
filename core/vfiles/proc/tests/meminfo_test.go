package vfiles

import (
	"fmt"
	vfiles "github.com/mikeyfennelly1/radharc/core/vfiles/proc"
	"testing"
)

func Test_Meminfo(t *testing.T) {
	t.Run("GetMemInfo", func(t *testing.T) {
		meminfo := vfiles.GetMemInfo()
		for _, entry := range meminfo {
			fmt.Println(entry)
		}
	})
}
