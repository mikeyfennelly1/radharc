package parse

import (
	"fmt"
	"strings"
)

type ConversionOperation struct {
	Apply func(string) (interface{}, error)
}

type KeyValPair struct {
	Key string
	Val string
}

func NewKeyVal(line string, separator string) (*KeyValPair, error) {
	// return an error if separator not in line
	if !strings.Contains(line, separator) {
		return nil, fmt.Errorf("KeyValSeparator %q not found in line: %q", separator, line)
	}

	// split line and return &KeyValPair for that line
	keyVal := strings.Split(line, separator)
	StrKVP := KeyValPair{
		Key: strings.TrimSpace(keyVal[0]),
		Val: strings.TrimSpace(keyVal[1]),
	}
	return &StrKVP, nil
}
