package parse

import (
	"errors"
	"strconv"
	"strings"
)

var (
	STR_TO_STR = ConvOperation[string]{
		function: strToStr,
	}
	STR_TO_INT = ConvOperation[int64]{
		function: strToInt,
	}
	STR_TO_FLOAT64 = ConvOperation[float64]{
		function: strToFloat64,
	}
	STR_TO_STR_SLICE = ConvOperation[[]string]{
		function: strToStrSlice,
	}
	YES_NO_TO_BOOL = ConvOperation[bool]{
		function: yesNoToBool,
	}
)

// return passed string
func strToStr(s string) (string, error) {
	return s, nil
}

// Parse string to int64
func strToInt(s string) (int64, error) {
	result, err := strconv.ParseInt(s, 10, 64)
	return result, err
}

// string to float64
func strToFloat64(s string) (float64, error) {
	result, err := strconv.ParseFloat(s, 64)
	return result, err
}

// Convert string to string slice
func strToStrSlice(s string) ([]string, error) {
	slice := strings.Split(s, " ")
	return slice, nil
}

// Convert yes/no to boolean value
func yesNoToBool(s string) (bool, error) {
	if s == "yes" {
		return true, nil
	} else if s == "no" {
		return false, nil
	} else {
		return false, errors.New("Input neither 'yes' or 'no'")
	}
}
