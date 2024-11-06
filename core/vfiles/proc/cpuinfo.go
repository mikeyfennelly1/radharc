package vfiles

import (
	parse "github.com/mikeyfennelly1/radharc/core/parse/kvp"
	"regexp"
	"strconv"
	"strings"
)

type cpuInfoValidTypes interface {
	int64 | string
}

func GetCPUInfo() parse.KVPParser {
	cpuinfoParser := parse.KVPParser{}
	parseToIntKeys := []string{"processor", "cpu family", "model", "stepping", "physical id", "siblings", "core id", "cpu cores", "apicid", "initial apicid", "cpuid level", "clflush size", "cache_alignment"}
	cpuinfoParser.AddKeyConvOperationPairs(parseToIntKeys, parse.StrToInt)
	parseToStrKeys := []string{"vendor_id", "model name", "microcode"}
	cpuinfoParser.AddKeyConvOperationPairs(parseToStrKeys, parse.StrToStr)
	popThreeValsReturnIntKeys := []string{"cache size"}
	cpuinfoParser.AddKeyConvOperationPairs(popThreeValsReturnIntKeys, parse.PopThreeCharsThenParseToInt)
	parseAddressSizesKeys := []string{"address sizes"}
	cpuinfoParser.AddKeyConvOperationPairs(parseAddressSizesKeys, parseAddressSizes)
	parseStrToStrSliceKeys := []string{"flags", "vmx flags", "bugs"}
	cpuinfoParser.AddKeyConvOperationPairs(parseStrToStrSliceKeys, parse.StrToStrSlice)

	return cpuinfoParser
}

type addressSizes struct {
	bitsPhysical int
	bitsVirtual  int
}

func parseAddressSizes(s string) (interface{}, error) {
	parsedAddressSizes := addressSizes{}

	// replace anything that isn't a digit with whitespace
	// split the values based on whitespace
	digitMatcher := regexp.MustCompile(`\D`)
	nonDigitRemoved := digitMatcher.ReplaceAllString(s, "")
	whitespaceMatcher := regexp.MustCompile(`\s`)
	// replace all whitespace with comma. Numbers are now separated as <num1>,<num2>
	whiteSpaceRemoved := whitespaceMatcher.ReplaceAllString(nonDigitRemoved, ",")
	splitOnComma := strings.Split(string(whiteSpaceRemoved), ",")

	bitsPhysical, err1 := strconv.Atoi(splitOnComma[0])
	bitsVirtual, err2 := strconv.Atoi(splitOnComma[0])
	if err1 != nil {
		parsedAddressSizes.bitsPhysical = bitsPhysical
	} else {
		return nil, err1
	}
	if err2 != nil {
		parsedAddressSizes.bitsVirtual = bitsVirtual
	} else {
		return nil, err2
	}

	return parsedAddressSizes, nil
}
