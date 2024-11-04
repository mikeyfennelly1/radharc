// contains utilities to read and parse files that exhibit key-value pair behavior
// @author Mikey Fennelly

package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type genericKVP[T any] struct {
	key   string
	value T
}

// Splits a key-value pair string at a splitter substring, and converts it to another
// type via the 'convert' function passed as an arg.
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

// convert a string to an integer
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
