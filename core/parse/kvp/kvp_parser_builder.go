package parse

type KeysAndCommonConvOp struct {
	Keys                      []string
	CommonConversionOperation func(s string) (interface{}, error)
}

func (parser *KVPParser) AddKeyConvOperationPair(key string, convOp func(s string) (interface{}, error)) {
	if parser == nil {
		panic("KVPParser is nil")
	}
	if parser.ConversionOpMap == nil {
		parser.ConversionOpMap = make(map[string]func(string) (interface{}, error))
	}
	parser.ConversionOpMap[key] = convOp
}
