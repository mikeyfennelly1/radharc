// contains utilities to read and parse files that exhibit key-value pair behavior
// @author Mikey Fennelly

package parse

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// struct for KVPParser - this is the core of the logic that serves as central point
// for parser logic and data
type KVPParser[T comparable] struct {
	ConversionOpMap map[string]func(string) (T, error)
	ParserRan       bool
	ResultMap       map[string]interface{}
}

func (parser *KVPParser[T]) RunParserOnFile(absoluteFilePath string, keyValueSeparator string) {
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

		separatorIndex := strings.Index(line, keyValueSeparator)
		if separatorIndex == -1 {
			continue
		}
		key := strings.TrimSpace(line[:separatorIndex])
		if convOperationValue, found := parser.ConversionOpMap[key]; found {
			stringVal := strings.TrimSpace(line[separatorIndex:])
			value, err := convOperationValue.function(stringVal)
			if err != nil {
				parser.ResultMap[key] = value
			}
			continue
		} else {
			continue
		}
	}
	parser.ParserRan = true
}
