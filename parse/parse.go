// Two modes: interpret and compile
package parse

import (
	"github.com/domluna/brainfuck-go/config"
	"github.com/domluna/brainfuck-go/lex"
	"github.com/domluna/brainfuck-go/program"
	"github.com/domluna/brainfuck-go/tape"
)

type Parser struct {
	lexer    *lex.Lexer
	conf     *config.Config
	prog     *program.Program
	fileName string
	currTok  lex.Token
	peekTok  lex.Token
	err      error
}

// New creates a new Parser.
// The Parser receives tokens from lexer and writes to out.
func New(fileName string, c *config.Config, l *lex.Lexer) *Parser {
	return &Parser{
		fileName: fileName,
		conf:     c,
		lexer:    l,
		prog:     program.NewProgram(tape.New(), c),
	}
}

// next returns the next token
func (p *Parser) next() lex.Token {
	tok := <-p.lexer.Tokens
	p.currTok = tok
	return tok
}

func (p *Parser) nextInst(tok lex.Token) program.Instruction {
	switch tok.Type {
	case lex.Ignore:
		// noop
	case lex.IncTape:
		return program.InstMoveHead{1}
	case lex.DecTape:
		return program.InstMoveHead{-1}
	case lex.IncByte:
		return program.InstAddToByte{1}
	case lex.DecByte:
		return program.InstAddToByte{-1}
	case lex.WriteByte:
		return program.InstWriteByte{}
	case lex.StoreByte:
		return program.InstSetByte{p.currTok.ByteVal}
	case lex.LoopEnter:
		return p.parseLoop()
	case lex.LoopExit:
		return nil
	}
	panic("parse: unreachable")
}

func (p *Parser) parseLoop() program.Instruction {
	loop := program.InstLoop{}
	insts := make([]program.Instruction, 0)
	for tok := p.next(); tok.Type != lex.EOF; tok = p.next() {
		i := p.nextInst(tok)
		if i == nil { // exit loop
			break
		}
		insts = append(insts, i)
	}
	return loop
}

func (p *Parser) Parse() (*program.Program, error) {
	for tok := p.next(); tok.Type != lex.EOF; tok = p.next() {
		i := p.nextInst(tok)
		p.prog.AddInst(i)
	}
	return p.prog, p.err
}
