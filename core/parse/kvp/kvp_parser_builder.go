package parse

type KeysAndCommonConvOp[T comparable] struct {
	Keys                      []string
	CommonConversionOperation func(string) (T, error)
}

func (parser *KVPParser[T]) AddKeyConvOperationPair(key string, convOp func(string) (T, error)) {
	parser.ConversionOpMap[key] = convOp
}

func (parser *KVPParser[T]) AddKeyConvOperationPairs(keysAndCommonConvOps KeysAndCommonConvOp[T]) {
	for _, key := range keysAndCommonConvOps.Keys {
		parser.AddKeyConvOperationPair(key, keysAndCommonConvOps.CommonConversionOperation)
	}
}
