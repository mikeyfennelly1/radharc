package parse

import (
	"errors"
	"strconv"
	"strings"
)

// return passed string
func strToStr(s string) (string, error) {
	return s, nil
}

var StrToStr = ConvOperation[string]{
	function: strToStr,
}

// Parse string to int64
func strToInt(s string) (int64, error) {
	result, err := strconv.ParseInt(s, 10, 64)
	return result, err
}

var StrToInt = ConvOperation[int64]{
	function: strToInt,
}

// string to float64
func strToFloat64(s string) (float64, error) {
	result, err := strconv.ParseFloat(s, 64)
	return result, err
}

var StrToFloat64 = ConvOperation[float64]{
	function: strToFloat64,
}

// Convert string to string slice
func strToStrSlice(s string) ([]string, error) {
	slice := strings.Split(s, " ")
	return slice, nil
}

var StrToStrSlice = ConvOperation[[]string]{
	function: strToStrSlice,
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
