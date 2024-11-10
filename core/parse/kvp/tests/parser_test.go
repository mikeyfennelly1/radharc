package parse

import (
	parse "github.com/mikeyfennelly1/radharc/core/parse/kvp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ParseLine(t *testing.T) {
	convOpMap := make(map[string]parse.ConversionOperation)
	convOpMap["processor"] = parse.ConversionOperation{parse.StrToInt}
	parser := parse.Parser{convOpMap}

	t.Run("<string> : <int64>", func(t *testing.T) {
		pKVP, err := parser.ParseLine("processor : 19", ":")
		got := (*pKVP)
		var expected int64 = 19
		assert.NoError(t, err)
		assert.NotNil(t, pKVP)
		// assert gotten val is as expected
		assert.Equal(t, expected, got.Value)
		// assert key is as expected
		assert.Equal(t, "processor", got.Key)
	})

	parser.ConvOpMap["vendor_id"] = parse.ConversionOperation{parse.StrToStr}
	t.Run("<string> : <string>", func(t *testing.T) {
		pKVP, err := parser.ParseLine("vendor_id : GenuineIntel", ":")
		got := (*pKVP)
		var expected string = "GenuineIntel"
		assert.NoError(t, err)
		assert.NotNil(t, pKVP)
		// assert gotten val is as expected
		assert.Equal(t, expected, got.Value)
		// assert key is as expected
		assert.Equal(t, "vendor_id", got.Key)
	})

	parser.ConvOpMap["arbitraryKey"] = parse.ConversionOperation{parse.StrToInt}
	t.Run("<string> : <nil>", func(t *testing.T) {
		pKVP, err := parser.ParseLine("arbitraryKey : ", ":")
		assert.Error(t, err)
		assert.Equal(t, "Nil value for key in line: \"arbitraryKey : \"", err.Error())
		assert.Nil(t, pKVP)
	})

	t.Run("<keyWithoutConversionOperation> : <arbitraryType>", func(t *testing.T) {
		pKVP, err := parser.ParseLine("keyWithoutConversionOperation : xxx", ":")
		assert.Error(t, err)
		assert.Equal(t, "ConversionOperation for key \"keyWithoutConversionOperation\", not in map", err.Error())
		assert.Nil(t, pKVP)
	})

	t.Run("no separator", func(t *testing.T) {
		pKVP, err := parser.ParseLine("a b", ":")
		assert.Error(t, err)
		assert.Nil(t, pKVP)
		assert.Equal(t, "KeyValSeparator \":\" not found in line: \"a b\"", err.Error())
	})
}
