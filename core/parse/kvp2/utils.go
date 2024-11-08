package parse

import (
	"errors"
	"strings"
)

type ConversionOperation struct {
	apply func(string) (interface{}, error)
}

type StringKeyVal struct {
	key string
	val string
}

func getStringKeyVal(line string, keyValSeparator string) (*StringKeyVal, error) {
	// return an error if separator not in line
	if !strings.Contains(line, keyValSeparator) {
		return nil, errors.New("KeyValSeparator not found in line: " + line)
	}

	// split line and return &StringKeyVal for that line
	keyVal := strings.Split(line, keyValSeparator)
	StrKVP := StringKeyVal{
		key: strings.TrimSpace(keyVal[0]),
		val: strings.TrimSpace(keyVal[1]),
	}
	return &StrKVP, nil
}

func findConversionOperationFromKey(key string, convOpMap map[string]ConversionOperation) (*ConversionOperation, error) {
	operation, ok := convOpMap[key]
	// if operation not in map return error
	if !ok {
		return nil, errors.New("ConversionOperation not found in map of ConversionOperations")
	}

	return &operation, nil
}
