package vfiles

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func GetAllRunningPids() []int64 {
	runningPids := make([]int64, 0)

	entries, err := ioutil.ReadDir("/proc/")
	if err != nil {
		fmt.Errorf("Err")
	}

	for _, entry := range entries {
		re := regexp.MustCompile(`^[0-9]+$`)
		if re.MatchString(entry.Name()) {
			pid, _ := strconv.ParseInt(entry.Name(), 10, 64)
			runningPids = append(runningPids, pid)
		}
	}

	return runningPids

}
