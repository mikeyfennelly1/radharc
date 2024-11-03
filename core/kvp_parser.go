package main

import (
	"fmt"
	"strconv"
	"strings"
)

const testString string = "key : value"

type ConversionOperation[T any] interface {
	operation(raw string) T
}

type genericKVP[T any] struct {
	key   string
	value T
}

func splitKVP[T any](input string, splitter string, convert func(string) (T, error)) (*genericKVP[T], error) {
	splitString := strings.Split(input, splitter)
	if len(splitString) < 2 {
		return nil, fmt.Errorf("input string: '%s' does not contain enough parts", input)
	}

	value, err := convert(strings.TrimSpace(splitString[1]))
	if err != nil {
		return nil, err
	}

	kvp := &genericKVP[T]{
		key:   strings.TrimSpace(splitString[0]),
		value: value,
	}

	return kvp, nil
}

func stringToInt(S string) (int, error) {
	return strconv.Atoi(S)
}

func main() {
	kvp, err := splitKVP("age : 30", ":", stringToInt)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return
	}
	fmt.Printf("Key: %s, Value: %d\n", kvp.key, kvp.value)
}
