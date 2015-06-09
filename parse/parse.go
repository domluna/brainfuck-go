// Two modes: interpret and compile
// TODO: start parsing those tokens!
// TODO: figure out how to compile.
// Current idea is to write a go program
// based on the tokens and then compile it
// then rm the program source.
package parse

import (
	"io"

	"github.com/domluna/brainfuck-go/config"
	"github.com/domluna/brainfuck-go/lex"
	"github.com/domluna/brainfuck-go/tape"
)

type Parser struct {
	name string
	l    *lex.Lexer
	conf config.Config
	tape tape.Tape
	out  io.Writer
}

// New creates a new Parser.
// The Parser receives tokens from lexer and writes to out.
func New(name string, conf *config.Config, lexer *lex.Lexer, out io.Writer) {
	return &Parser{
		name: name,
		conf: conf,
		l:    lexer,
		tape: tape.New(300000),
		out:  out,
	}
}
