// Brainfuck program.
//
// A Brainfuck Program consists of a list of Instructions
// executed sequentially.
//
// Instructions are evaluated on a Tape.
package program

import (
	"fmt"

	"github.com/domluna/brainfuck-go/config"
)

type Program struct {
	insts []Instruction
	t     Tape
	conf  *config.Config
}

func NewProgram(t Tape, c *config.Config) *Program {
	return &Program{
		insts: make([]Instruction, 0),
		t:     t,
		conf:  c,
	}
}

func (p *Program) String() string {
	s := ""
	for _, i := range p.insts {
		s += fmt.Sprintf("%s\n", i)
	}
	return s
}

// AddInst adds i to the program Instruction list.
// Instructions are executed sequentially in the order
// they are added.
func (p *Program) AddInst(i Instruction) {
	p.insts = append(p.insts, i)
}

// Run evaluates the Program's instructions.
func (p *Program) Run() []byte {
	for _, i := range p.insts {
		i.Eval(p.t)
	}
	return p.t.Output()
}

// TODO:
// Compile compiles the Program to a standalone binary executable.
func (p *Program) Compile() {
	panic("program: compile not implemented")
}
