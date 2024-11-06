package vfiles

import (
	parse "github.com/mikeyfennelly1/radharc/core/parse/kvp"
)

type cpuInfoValidTypes interface {
	int64 | string
}

func GetCPUInfo() parse.KVPParser {
	cpuinfoParser := parse.KVPParser{}
	cpuinfoParser.AddKeyConvOperationPair("processor", parse.StrToStr)
	return cpuinfoParser
}
