package program_test

import (
	"reflect"
	"testing"

	"github.com/domluna/brainfuck-go/program"
)

var tests = []struct {
	in  []program.Instruction
	out []program.Instruction
}{
	// empty
	{
		in:  []program.Instruction{},
		out: []program.Instruction{},
	},
	// only consecutive
	{
		in: []program.Instruction{
			program.InstMoveHead{1},
			program.InstMoveHead{1},
			program.InstMoveHead{1},
			program.InstMoveHead{1},
			program.InstMoveHead{1},
		},
		out: []program.Instruction{
			program.InstMoveHead{5},
		},
	},
	// interleaving
	{
		in: []program.Instruction{
			program.InstMoveHead{1},
			program.InstMoveHead{1},
			program.InstWriteToOutput{},
			program.InstAddToByte{1},
			program.InstAddToByte{-2},
			program.InstWriteToOutput{},
			program.InstMoveHead{-1},
			program.InstAddToByte{10},
		},
		out: []program.Instruction{
			program.InstMoveHead{2},
			program.InstWriteToOutput{},
			program.InstAddToByte{-1},
			program.InstWriteToOutput{},
			program.InstMoveHead{-1},
			program.InstAddToByte{10},
		},
	},
	// loop
	{
		in: []program.Instruction{
			program.InstMoveHead{1},
			program.InstMoveHead{1},
			program.InstWriteToOutput{},
			program.InstReadFromInput{},
			program.InstLoop{
				[]program.Instruction{
					program.InstMoveHead{1},
					program.InstMoveHead{1},
					program.InstWriteToOutput{},
					program.InstReadFromInput{},
				},
			},
			program.InstMoveHead{-1},
		},
		out: []program.Instruction{
			program.InstMoveHead{2},
			program.InstWriteToOutput{},
			program.InstReadFromInput{},
			program.InstLoop{
				[]program.Instruction{
					program.InstMoveHead{2},
					program.InstWriteToOutput{},
					program.InstReadFromInput{},
				},
			},
			program.InstMoveHead{-1},
		},
	},
}

func Test_Optimize(t *testing.T) {
	for i, tt := range tests {
		res := program.Optimize(tt.in)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("Test(%d): expected %v, got %v", i, tt.out, res)
		}
	}
}
