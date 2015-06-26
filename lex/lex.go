// The structure of this lexer is derived from
// Rob Pike's APL interpreter "ivy" (github.com/robpike/ivy).
//
// Check it out, it's awesome!
package lex

import (
	"fmt"
	"io"

	"github.com/domluna/brainfuck-go/config"
)

//go:generate stringer -type=Type
type Type int

// EOF being the zero value has nice implications
// for when we close our Token channel down the line.
const (
	EOF Type = iota // zero value
	NewLine
	IncTape   // '>' increment tape head
	DecTape   // '<' decrement tape head
	IncByte   // '+' increment byte value at tape head
	DecByte   // '-' decrement byte value at tape head
	WriteByte // '.' write byte to output
	StoreByte // ',' store byte from input to tape header
	LoopEnter // '['
	LoopExit  // ']'
)

// Token represents a tokenized byte.
type Token struct {
	Type    Type
	Line    int
	Pos     int
	ByteVal byte
}

func (t Token) String() string {
	switch t.Type {
	case EOF:
		return "EOF"
	}
	return fmt.Sprintf("%s: %q", t.Type, t.ByteVal)
}

const eof = -1

type stateFn func(*Lexer) stateFn

// Lexer reads an input byte stream and generates
// Brainfuck Tokens.
type Lexer struct {
	Tokens   chan Token     // lexed tokens
	r        io.ByteReader  // input stream
	conf     *config.Config //
	fileName string         // name of file being lexed
	lineNo   int            // line number
	pos      int            // position in the line
	input    byte           // current character
	done     bool
	state    stateFn
}

// New creates a Lexer. The Lexer reads from r and outputs
// Token values to its channel Tokens.
//
// The Lexer works concurrently.
func New(fileName string, c *config.Config, r io.ByteReader) *Lexer {
	l := &Lexer{
		Tokens:   make(chan Token),
		conf:     c,
		r:        r,
		lineNo:   1,
		pos:      0,
		fileName: fileName,
	}
	go l.run()
	return l
}

// Pos returns the line position of the lexer.
func (l *Lexer) Pos() int {
	return l.pos
}

// Line returns the line number the lexer is on.
func (l *Lexer) Line() int {
	return l.lineNo
}

func (l *Lexer) run() {
	for l.state = lexMain; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.Tokens)
}

// we return a rune here so we can use -1 as a value
// byte ranges from 0..255
func (l *Lexer) next() rune {
	c, err := l.r.ReadByte()

	if err != nil { // EOF
		l.done = true
		return eof
	}

	l.input = c
	l.pos++
	return rune(c)
}

// send the Token for Type t through the Tokens channel.
func (l *Lexer) send(t Type) {
	if t == NewLine {
		l.lineNo++
		l.pos = 1
	}

	tok := Token{
		Type:    t,
		ByteVal: l.input,
		Line:    l.lineNo,
		Pos:     l.pos,
	}

	l.conf.Debug("lex: <%q %d:%d> sending: %s\n", l.fileName, l.lineNo, l.pos, tok)
	l.Tokens <- tok
}

// Start the lexer by choosing the initial state
// in this state the lexer chooses between all states.
// Since Brainfuck is so simple there's no point in making
// more states, but in a more complicated language we would
// want to return different states here.
func lexMain(l *Lexer) stateFn {
	r := l.next()
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
		l.send(WriteByte)
	case ',':
		l.send(StoreByte)
	case '[':
		l.send(LoopEnter)
	case ']':
		l.send(LoopExit)
	}

	return lexMain
}
