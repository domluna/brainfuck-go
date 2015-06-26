package program

import (
	"fmt"
	"io"
)

// Instructions represents a Brainfuck instruction.
type Instruction interface {
	// Evaluate the instruction on the Tape.
	Eval(t Tape, in io.ByteReader, out io.ByteWriter)

	// String representation of the instruction.
	String() string
}

//
// Instructions
//

// InstMoveHead moves the Tape head V spaces.
type InstMoveHead struct {
	V int
}

func (i InstMoveHead) Eval(t Tape, in io.ByteReader, out io.ByteWriter) {
	t.MoveHead(i.V)
}

func (i InstMoveHead) String() string {
	return fmt.Sprintf("InstMoveHead{%d}", i.V)
}

// InstAddToByte adds V to current byte value at the Tape head.
// V is an integer representation of the byte value to be added.
type InstAddToByte struct {
	V int
}

func (i InstAddToByte) Eval(t Tape, in io.ByteReader, out io.ByteWriter) {
	t.AddToByte(i.V)
}

func (i InstAddToByte) String() string {
	return fmt.Sprintf("InstAddToByte{%d}", i.V)
}

// InstWriteToOutput writes the byte at the Tape head to output.
// The output is an io.ByteWriter.
type InstWriteToOutput struct{}

func (i InstWriteToOutput) Eval(t Tape, in io.ByteReader, out io.ByteWriter) {
	b := t.GetByte()
	out.WriteByte(b)
}

func (i InstWriteToOutput) String() string {
	return "InstWriteToOutput"
}

// InstReadFromInput reads a byte from the input. The byte is then written
// to the slot at the Tape head.
//
// input is an io.ByteReader.
type InstReadFromInput struct{}

func (i InstReadFromInput) Eval(t Tape, in io.ByteReader, out io.ByteWriter) {
	b, _ := in.ReadByte()
	if b == byte(0) {
		return
	}
	t.SetByte(b)
}

func (i InstReadFromInput) String() string {
	return "InstReadFromInput"
}

// InstLoop loops over an Instruction slice. Repeats the loop
// until the exit condition is met.
//
// The exit condition is at the start and end of each loop check
// if the byte at the Tape head == byte(0). If true, exit.
type InstLoop struct {
	Insts []Instruction
}

func (i InstLoop) Eval(t Tape, in io.ByteReader, out io.ByteWriter) {
	for {
		// loop exit condition
		if t.GetByte() == byte(0) {
			break
		}

		for _, ii := range i.Insts {
			ii.Eval(t, in, out)
		}

		// loop exit condition
		if t.GetByte() == byte(0) {
			break
		}

	}
}

func (i InstLoop) String() string {
	s := "InstLoop\n"
	for _, ii := range i.Insts {
		s += fmt.Sprintf("%s\n", ii)
	}
	return s
}
