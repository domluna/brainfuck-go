package lex

import (
	"reflect"
	"strings"
	"testing"
)

var tests = []struct {
	in  string
	out []Type
}{
	{
		"><+-.,[]",
		[]Type{IncTape, DecTape, IncByte, DecByte,
			WriteByte, StoreByte, LoopEnter, LoopExit},
	},
	{
		"+\nabcd\n-",
		[]Type{IncByte, NewLine, NewLine, DecByte},
	},
}

func TestLexer(t *testing.T) {

	for _, tt := range tests {
		l := New("", nil, strings.NewReader(tt.in))
		typs := make([]Type, 0)

		for v := range l.Tokens {
			typs = append(typs, v.Type)
		}

		if !reflect.DeepEqual(tt.out, typs) {
			t.Errorf("expected %v, got %v", tt.out, typs)
		}
	}
}
