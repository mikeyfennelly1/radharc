package parse

func (parser *Parser) AddConvOps(keys []string, convOp ConversionOperation) {
	for _, key := range keys {
		parser.ConvOpMap[key] = convOp
	}
}
