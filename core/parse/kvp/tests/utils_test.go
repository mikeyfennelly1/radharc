package parse

import (
	parse "github.com/mikeyfennelly1/radharc/core/parse/kvp2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewKeyVal(t *testing.T) {
	t.Run("Perfect case (key, value and separator exist) - no spaces", func(t *testing.T) {
		kvp, err := parse.NewKeyVal("a:b", ":")
		assert.NoError(t, err)
		assert.NotNil(t, kvp)
		assert.Equal(t, "a", kvp.Key)
		assert.Equal(t, "b", kvp.Val)
	})

	t.Run("Perfect case (key, value and separator exist) - has spaces", func(t *testing.T) {
		kvp, err := parse.NewKeyVal(" a : b ", ":")
		assert.NoError(t, err)
		assert.NotNil(t, kvp)
		assert.Equal(t, "a", kvp.Key)
		assert.Equal(t, "b", kvp.Val)
	})

	t.Run("wrong separator", func(t *testing.T) {
		kvp, err := parse.NewKeyVal(" a : b ", "#")
		assert.Error(t, err)
		assert.Equal(t, "KeyValSeparator \"#\" not found in line: \" a : b \"", err.Error())
		assert.Nil(t, kvp)
	})
}
