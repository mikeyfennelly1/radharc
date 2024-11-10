package parse

import (
	parse "github.com/mikeyfennelly1/radharc/core/parse/kvp2"
	"testing"
)

func TestGetStringKeyVal(t *testing.T) {
	tests := []struct {
		name            string
		line            string
		keyValSeparator string
		want            *parse.StringKeyVal
	}{
		{"Perfect case (key, value and separator exist)",
			"key : value",
			":",
			&parse.StringKeyVal{Key: "key", Val: "value"},
		},
		{
			"No value",
			"key : ",
			":",
			&parse.StringKeyVal{Key: "key", Val: ""},
		},
		{
			"No key",
			" : value",
			":",
			&parse.StringKeyVal{Key: "", Val: "value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := parse.GetStringKeyVal(tt.line, tt.keyValSeparator)
			if got.Val != tt.want.Val || got.Key != tt.want.Key {
				t.Errorf("%s: parse.GetStringKeyVal(\"%s\", \"%s\") => %s; want: %s\n", tt.name, tt.line, tt.keyValSeparator, got, tt.want)
			}
		})
	}
}
