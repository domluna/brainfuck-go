package program

// Tape represents a tape a Brainfuck program executes on.
type Tape interface {

	// Moves the tape head i spots forward or backward,
	// depending on i.
	MoveHead(i int)

	// Adds i to the byte value at the tape head.
	// This should deal with over/under flows by wrapping.
	AddToByte(i int)

	// Set the byte at the tape head to b.
	SetByte(b byte)

	// Return the byte at the tape head.
	GetByte() byte

	// Get the position of the tape head.
	GetHead() int

	// Set the position of the tape head to i.
	SetHead(i int)
}
