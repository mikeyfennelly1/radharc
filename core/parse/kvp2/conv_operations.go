package parse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// return passed string
func StrToStr(s string) (interface{}, error) {
	return s, nil
}

// Parse string to int64
func StrToInt(s string) (interface{}, error) {
	result, err := strconv.ParseInt(s, 10, 64)
	return result, err
}

// string to float64
func StrToFloat64(s string) (interface{}, error) {
	result, err := strconv.ParseFloat(s, 64)
	return result, err
}

// Convert string to string slice
func StrToStrSlice(s string) (interface{}, error) {
	slice := strings.Split(s, " ")
	return slice, nil
}

// Convert yes/no to boolean value
func YesNoToBool(s string) (interface{}, error) {
	if s == "yes" {
		return true, nil
	} else if s == "no" {
		return false, nil
	} else {
		return false, errors.New("Input neither 'yes' or 'no'")
	}
}

func PopThreeCharsThenParseToInt(s string) (interface{}, error) {
	if len(s) > 3 {
		returnStr := s[:len(s)-3]
		return returnStr, nil
	} else {
		return nil, fmt.Errorf("lngth of inputted string '%s' is <= 3", s)
	}
}
