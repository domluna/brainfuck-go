// Two modes: interpret and compile
package parse

import (
	"github.com/domluna/brainfuck-go/config"
	"github.com/domluna/brainfuck-go/lex"
	"github.com/domluna/brainfuck-go/program"
)

type Parser struct {
	lexer    *lex.Lexer
	conf     *config.Config
	fileName string
	currTok  lex.Token
	err      error
}

// New creates a new Parser.
// The Parser receives tokens from lexer and writes to out.
func New(fileName string, c *config.Config, l *lex.Lexer) *Parser {
	return &Parser{
		fileName: fileName,
		conf:     c,
		lexer:    l,
	}
}

// next returns the next token
func (p *Parser) next() lex.Token {
	tok := <-p.lexer.Tokens
	for tok.Type == lex.NewLine {
		// inc newline
		tok = <-p.lexer.Tokens
	}
	p.currTok = tok
	return tok
}

func (p *Parser) nextInst(tok lex.Token) program.Instruction {
	switch tok.Type {
	case lex.IncTape:
		return program.InstMoveHead{1}
	case lex.DecTape:
		return program.InstMoveHead{-1}
	case lex.IncByte:
		return program.InstAddToByte{1}
	case lex.DecByte:
		return program.InstAddToByte{-1}
	case lex.WriteByte:
		return program.InstWriteToOutput{}
	case lex.StoreByte:
		return program.InstReadFromInput{}
	case lex.LoopEnter:
		return p.parseLoop()
	case lex.LoopExit:
		return nil
	}
	panic("parse: unreachable")
}

func (p *Parser) parseLoop() program.Instruction {
	insts := make([]program.Instruction, 0)
	for tok := p.next(); tok.Type != lex.EOF; tok = p.next() {
		i := p.nextInst(tok)
		if i == nil { // exit loop
			break
		}
		insts = append(insts, i)
	}
	return program.InstLoop{insts}
}

func (p *Parser) Parse() ([]program.Instruction, error) {
	prog := make([]program.Instruction, 0)
	for tok := p.next(); tok.Type != lex.EOF; tok = p.next() {
		i := p.nextInst(tok)
		// p.conf.Debug("parse: <%s %d:%d> adding Instruction: %s\n", p.fileName,
		// p.lexer.Line(), p.lexer.Pos(), i)
		prog = append(prog, i)
		// prog.AddInst(i)
	}
	return prog, p.err
}
