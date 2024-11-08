package parse

type KVP struct {
	key   string
	value interface{}
}

type Parser struct {
	convOpMap map[string]ConversionOperation
}

func (parser *Parser) parseLine(line string, keyValSeparator string) (*KVP, error) {
	// get key and value as string types
	kvpString, lineNotFoundErr := getStringKeyVal(line, keyValSeparator)
	if lineNotFoundErr != nil {
		return nil, lineNotFoundErr
	}

	// find ConversionOperation for the key
	key := (*kvpString).key
	convOp, convOpNotFoundErr := findConversionOperationFromKey(key, parser.convOpMap)
	if convOpNotFoundErr != nil {
		return nil, convOpNotFoundErr
	}

	// apply conversion operation to the string value
	valueString := (*kvpString).val
	convertedValue, unableToConvertErr := convOp.apply(valueString)
	if unableToConvertErr != nil {
		return nil, unableToConvertErr
	}

	// create struct for data and return
	keyValuePair := KVP{key: key, value: convertedValue}
	return &keyValuePair, nil
}
