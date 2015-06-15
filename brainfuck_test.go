package main

import (
	"strings"
	"testing"

	"github.com/domluna/brainfuck-go/config"
	"github.com/domluna/brainfuck-go/lex"
	"github.com/domluna/brainfuck-go/parse"
)

const helloWorldProg = `
++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.
`

var tests = []struct {
	name string
	in   string
	out  string
}{
	{
		"hello_world.b",
		helloWorldProg,
		"Hello World!\n",
	},
}

func TestBrainfuck(t *testing.T) {

	var conf *config.Config
	conf = config.New(true)

	for _, tt := range tests {
		lexer := lex.New(tt.name, conf, strings.NewReader(tt.in))
		parser := parse.New(tt.name, conf, lexer)
		prog, err := parser.Parse()
		if err != nil {
			t.Fatalf("expected <nil>, got %q", err)
		}

		result := prog.Run()
		if string(result) != tt.out {
			t.Errorf("expected %s, got %s", tt.out, string(result))
		}

	}
}
