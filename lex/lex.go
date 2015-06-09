package lex

//go:generate stringer -type Type

import (
	"fmt"
	"io"
	"log"

	"github.com/domluna/brainfuck-go/config"
)

// Token represents a tokenized string.
type Token struct {
	Type Type
	Text string
	Line int
	Pos  int
}

type Type int

const (
	EOF   Type = iota // zero value
	Error             // error occured during reading
	Newline
	IncTape   // '>' increment tape position
	DecTape   // '<' increment tape position
	IncByte   // '+' increment byte value at tape position
	DecByte   // '-' decrement byte value at tape position
	ReadByte  // '.' read the byte at tape position
	StoreByte // ',' store the byte at tape position
	StartLoop // '['
	EndLoop   // ']'
)

func (t Token) String() string {
	switch t.Type {
	case EOF:
		return "EOF"
	case ERROR:
		return fmt.Sprintf("ERROR: %s", t.Text)
	}
	return fmt.Sprintf("%s: %q", t.Type, t.Text)
}

const eof = -1

type stateFn func(*Lexer) stateFn

// Lexer reads an input byte stream and generates
// Brainfuck Tokens.
type Lexer struct {
	Tokens chan Token     // lexed tokens
	r      io.ByteReader  // input stream
	conf   *config.Config //
	name   string         // name of file being lexed
	line   int            // line number
	pos    int            // position in the line
	input  string         // current character
	done   bool
	state  stateFn
}

// New creates a Lexer. The Lexer reads from r and outputs
// Token values to its channel Tokens.
func New(name string, conf *config.Config, r io.ByteReader) *Lexer {
	l := &Lexer{
		Tokens: make(chan Token),
		r:      r,
		line:   1,
		pos:    1,
		name:   name,
	}
	go l.run()
	return l
}

func (l *Lexer) run() {
	for l.state = lexMain(l); l.state != nil; {
		l.state = l.state(l)
	}
	close(l.Tokens)
}

func (l *Lexer) next() rune {
	c, err := l.r.ReadByte()
	if err != nil { // EOF
		l.done = true
		return eof
	}

	l.input = string(c)
	l.pos++
	return rune(c)
}

// send the Token for Type t through the Token channel.
func (l *Lexer) send(t Type) {
	if t == Newline {
		l.line++
		l.pos = 1
	}

	tok := Token{
		Type: t,
		Text: l.input,
		Line: t.line,
		Pos:  t.pos,
	}

	if l.conf.Debug() {
		log.Printf("%s %d:%d sending: %s\n", l.name, l.line, l.pos, tok)
	}

	l.Tokens <- tok
}

// Start the lexer by choosing the initial state
// in this state the lexer chooses between all states.
// Since Brainfuck is so simple there's no point in making
// more states, but in a more complicated language we would
// want to return different states here.
func lexMain(l *Lexer) stateFn {
	r := l.read()
	switch r {
	case eof: // This will terminate the loop in l.run()
		return nil
	case '\n':
		l.send(NewLine)
	case '>':
		l.send(IncTape)
	case '<':
		l.send(DecTape)
	case '+':
		l.send(IncByte)
	case '-':
		l.send(DecByte)
	case '.':
		l.send(ReadByte)
	case ',':
		l.send(StoreByte)
	case '[':
		l.send(StartLoop)
	case ']':
		l.send(EndLoop)
	default: // ERROR
		l.send(ERROR)
	}

	return lexMain
}
