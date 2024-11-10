package parse

import (
	"fmt"
)

type KVP struct {
	Key   string
	Value interface{}
}

type Parser struct {
	ConvOpMap map[string]ConversionOperation
}

func (parser *Parser) ParseLine(line string, keyValSeparator string) (*KVP, error) {
	// get key and value as string types
	pKVPString, lineNotFoundErr := NewKeyVal(line, keyValSeparator)
	if lineNotFoundErr != nil {
		return nil, lineNotFoundErr
	} else if pKVPString.Val == "" {
		return nil, fmt.Errorf("Nil value for key in line: %q", line)
	}

	// find ConversionOperation for the key
	key := (*pKVPString).Key
	convOp, ok := parser.ConvOpMap[key]
	if !ok {
		return nil, fmt.Errorf("ConversionOperation for key %q, not in map", key)
	}

	// apply conversion operation to the string value
	valueString := (*pKVPString).Val
	convertedValue, unableToConvertErr := convOp.Apply(valueString)
	if unableToConvertErr != nil {
		return nil, unableToConvertErr
	}

	// create struct for data and return
	keyValuePair := KVP{Key: key, Value: convertedValue}
	return &keyValuePair, nil
}
