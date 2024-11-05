// contains utilities to read and parse files that exhibit key-value pair behavior
// @author Mikey Fennelly

package parse

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Converter interface {
	Convert(string) (any, error)
}

type convOperation[T any] struct {
	function func(string) (T, error)
}

func (c convOperation[T]) Convert(s string) (any, error) {
	return c.function(s)
}

var conv_int64 = func(s string) (int64, error) {
	parsedInt, err := strconv.ParseInt(s, 10, 64)
	return parsedInt, err
}

type KVP[T any] struct {
	key   string
	value T
}

// Splits a key-value pair string at a splitter substring, and converts it to another
// type via the 'convert' function passed as an arg.
func lineParser[T any](input string, splitter string, convert func(string) (T, error)) (*KVP[T], error) {
	splitString := strings.Split(input, splitter)
	if len(splitString) < 2 {
		return nil, fmt.Errorf("input string: '%s' does not contain enough parts", input)
	}

	value, err := convert(strings.TrimSpace(splitString[1]))
	if err != nil {
		return nil, err
	}

	kvp := &KVP[T]{
		key:   strings.TrimSpace(splitString[0]),
		value: value,
	}

	return kvp, nil
}

type KVPParser struct {
	opMap     map[string]convOperation[interface{}]
	parserRan bool
	resultMap map[string]interface{}
}

func (cm *KVPParser) AddConvOperation(key string, convOp convOperation[interface{}]) {
	cm.opMap[key] = convOp
}

func (parser *KVPParser) Parse(absoluteFilePath string, separator string) {
	file, err := os.Open(absoluteFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		separatorIndex := strings.Index(line, separator)
		if separatorIndex == -1 {
			continue
		}
		key := strings.TrimSpace(line[:separatorIndex])
		if convOperationValue, found := parser.opMap[key]; found {
			stringVal := strings.TrimSpace(line[separatorIndex:])
			value, err := convOperationValue.function(stringVal)
			if err != nil {
				parser.resultMap[key] = value
			}
			continue
		} else {
			continue
		}
	}
	parser.parserRan = true
}
