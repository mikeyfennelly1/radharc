package main

import (
	"fmt"
	vfiles "github.com/mikeyfennelly1/radharc/core/vfiles/proc"
)

func main() {
	cpuinfoParser := vfiles.GetCPUInfo()
	fmt.Println(cpuinfoParser)
}
