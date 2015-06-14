package program

import "fmt"

type Instruction interface {
	// Evaluate the instruction on the Tape.
	Eval(Tape)

	// String representation of the instruction.
	String() string
}

//
// Instructions
//

type InstMoveHead struct {
	V int
}

func (i InstMoveHead) String() string {
	return fmt.Sprintf("InstMoveHead %d", i.V)
}

func (i InstMoveHead) Eval(t Tape) {
	t.MoveHead(i.V)
}

type InstAddToByte struct {
	V int
}

func (i InstAddToByte) String() string {
	return fmt.Sprintf("InstAddToByte %d", i.V)
}

func (i InstAddToByte) Eval(t Tape) {
	t.AddToByte(i.V)
}

type InstWriteByte struct{}

func (i InstWriteByte) String() string {
	return fmt.Sprintf("InstWriteByte")
}

func (i InstWriteByte) Eval(t Tape) {
	t.WriteByte()
}

type InstSetByte struct {
	B byte
}

func (i InstSetByte) String() string {
	return fmt.Sprintf("InstSetByte %q, %v", i.B, i.B)
}

func (i InstSetByte) Eval(t Tape) {
	t.SetByte(i.B)
}

type InstLoop struct {
	Insts []Instruction
}

func (i InstLoop) String() string {
	s := "InstLoop\n"
	for _, ii := range i.Insts {
		s += fmt.Sprintf("  %s\n", ii)
	}
	return s
}

func (i InstLoop) Eval(t Tape) {
	headStart := t.GetHead()
	for {
		for _, ii := range i.Insts {
			ii.Eval(t)
		}

		// loop exit condition
		if t.GetByte() != byte(0) {
			return
		}

		// Reset tape head
		t.SetHead(headStart)
	}
}
