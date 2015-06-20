package tape

type Tape struct {
	head int // position of the head in the tape
	tape []byte
}

func New() *Tape {
	return &Tape{
		head: 0,
		tape: make([]byte, 1),
	}
}

// MoveHead moves the tape head from its current position
// i spots forward/backward.
func (t *Tape) MoveHead(i int) {
	if t.head+i < 0 {
		panic("tape: out of bounds, cannot have a negative tape header")
	}

	if t.head+i+1 > len(t.tape) {
		t.tape = append(t.tape, make([]byte, i+1)...)
	}
	t.head += i
}

// AddToByte adds i to the byte value at the
// tape head. Deals with over/under flow issues.
//
// Ex: if the current value is 0 and we add -10, we'll get
// 246.
func (t *Tape) AddToByte(i int) {
	cb := t.tape[t.head]
	if (int(cb)+i)%256 < 0 {
		t.tape[t.head] = byte(int(cb) + (i % 256) + 256)
		return
	}
	t.tape[t.head] = byte((int(cb) + i) % 256)
}

// GetByte returns value of the byte at the tape head.
func (t *Tape) GetByte() byte {
	b := t.tape[t.head]
	return b
}

// SetByte stores b at the tape head.
func (t *Tape) SetByte(b byte) {
	t.tape[t.head] = b
}

// GetHead returns the position of the tape head.
func (t *Tape) GetHead() int {
	return t.head
}

// SetHead sets the tape head to to i.
func (t *Tape) SetHead(i int) {
	if i < 0 || i > len(t.tape) {
		panic("tape: cannot set head to negative index")
	}
	t.head = i
}
