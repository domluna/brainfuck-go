package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/domluna/brainfuck-go/config"
	"github.com/domluna/brainfuck-go/lex"
	"github.com/domluna/brainfuck-go/parse"
)

var tests = []struct {
	name string
	in   string
	out  string
}{
	{
		"hello_world.b",
		"++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
		"Hello world!\n",
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

		buf := prog.Run()
		fmt.Println(buf)

		if string(buf) != tt.out {
			t.Errorf("expected %s, got %s", tt.out, string(buf))
		}

	}
}
