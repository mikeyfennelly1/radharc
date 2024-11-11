package vfiles

import (
	"bufio"
	"fmt"
	parse "github.com/mikeyfennelly1/radharc/core/parse/kvp"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cpuInfoValidTypes interface {
	int64 | string
}

func GetCPUInfo() []map[string]interface{} {
	parser := parse.Parser{ConvOpMap: make(map[string]parse.ConversionOperation)}

	strToStrKeys := []string{"vendor_id", "model name", "microcode"}
	parser.AddConvOps(strToStrKeys, parse.ConversionOperation{parse.StrToStr})

	strToStrListKeys := []string{"flags", "vmx flags", "bugs"}
	parser.AddConvOps(strToStrListKeys, parse.ConversionOperation{parse.StrToStrSlice})

	strToIntKeys := []string{"processor", "cpu family", "model", "stepping", "physical id", "siblings", "core id", "cpu cores", "apicid", "initial apicid", "cpuid level", "clflush size", "cache_alignment"}
	parser.AddConvOps(strToIntKeys, parse.ConversionOperation{parse.StrToInt})

	strToFloatKeys := []string{"cpu MHz", "bogomips"}
	parser.AddConvOps(strToFloatKeys, parse.ConversionOperation{parse.StrToFloat64})

	addressSizesKeys := []string{"address sizes"}
	parser.AddConvOps(addressSizesKeys, parse.ConversionOperation{parseAddressSizes})

	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Errorf("Error opening the file /proc/cpuinfo")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cpuInfoSections := make([]map[string]interface{}, 0)
	thisSectionKvps := make(map[string]interface{})
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			cpuInfoSections = append(cpuInfoSections, thisSectionKvps)
		}

		thisLineKVP, err := parser.ParseLine(line, ":")
		if err != nil {
			continue
		}
		thisSectionKvps[thisLineKVP.Key] = thisLineKVP.Value
	}

	return cpuInfoSections
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
