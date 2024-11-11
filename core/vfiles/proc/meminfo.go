package vfiles

import (
	"bufio"
	"fmt"
	parse "github.com/mikeyfennelly1/radharc/core/parse/kvp"
	"os"
)

// returns a memory info in KB, or
func GetMemInfo() map[string]interface{} {
	parser := parse.Parser{ConvOpMap: make(map[string]parse.ConversionOperation)}

	kilobyteValKeys := []string{"MemTotal", "MemFree", "MemAvailable", "Buffers", "Cached", "SwapCached", "Active", "Inactive", "Active(anon)", "Inactive(anon)", "Active(file)", "Inactive(file)", "Unevictable", "Mlocked", "SwapTotal", "SwapFree", "Zswap", "Zswapped", "Dirty", "Writeback", "AnonPages", "Mapped", "Shmem", "KReclaimable", "Slab", "SReclaimable", "SUnreclaim", "KernelStack", "PageTables", "SecPageTables", "NFS_Unstable", "Bounce", "WritebackTmp", "CommitLimit", "Committed_AS", "VmallocTotal", "VmallocUsed", "VmallocChunk", "Percpu", "HardwareCorrupted", "AnonHugePages", "ShmemHugePages", "ShmemPmdMapped", "FileHugePages", "FilePmdMapped", "Unaccepted", "Hugepagesize", "Hugetlb", "DirectMap4k", "DirectMap2M", "DirectMap1G"}
	parser.AddConvOps(kilobyteValKeys, parse.ConversionOperation{parse.PopThreeCharsThenParseToInt})

	intValKeys := []string{"HugePages_Total", "HugePages_Free", "HugePages_Rsvd", "HugePages_Surp"}
	parser.AddConvOps(intValKeys, parse.ConversionOperation{parse.StrToInt})

	file, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Errorf("Error opening the file /proc/meminfo")
	}
	defer file.Close()

	meminfo := make(map[string]interface{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		thisLineKVP, err := parser.ParseLine(line, ":")
		if err != nil {
			continue
		}
		meminfo[thisLineKVP.Key] = thisLineKVP.Value
	}
	return meminfo
}
