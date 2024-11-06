package parse

func getConvOperation[T any](def T) func(string) (interface{}, error) {
	v := any(def)
	switch v.(type) {
	case int64:
		return StrToInt
	case string:
		return StrToStr
	case float64:
		return StrToFloat64
	case []string:
		return StrToStrSlice
	case bool:
		return YesNoToBool
	default:
		return nil
	}
}
