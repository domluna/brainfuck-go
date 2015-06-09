package tape

// Used for over/under flow checking
const (
	upBound = byte(255) // highest byte value
	loBound = byte(0)   // lowest byte value
)

type Tape struct {
	pos  int // position of the header in the tape
	tape []byte
}

func New(n int) *Tape {
	return &Tape{
		pos:  0,
		tape: make(tape, n, n),
	}
}

// IncPos increments the position of t.pos by 1.
func (t *Tape) IncPos() {
	t.pos++
}

// DecPos decrements the position of t.pos by 1.
func (t *Tape) DecPos() {
	t.pos--
}

// IncByte increments the byte value at t.pos by 1.
// If the byte value is equal to upBound it will
// wrap and become loBound.
func (t *Tape) IncByte() {
	cb := t.tape[t.pos]
	if cb == upBound {
		t.tape[t.pos] = loBound
		return
	}
	t.tape[t.pos] = cb + 1
}

// DecByte decrements the byte value at t.pos by 1.
// If the byte value is equal to loBound it will
// wrap and become upBound.
func (t *Tape) DecByte() {
	cb := t.tape[t.pos]
	if cb == loBound {
		t.tape[t.pos] = upBound
		return
	}
	t.tape[t.pos] = cb - 1
}

// ReadByte reads the byte at t.pos and returns it.
func (t *Tape) ReadByte() byte {
	return t.tape[t.pos]
}

// StoreByte stores b at the current Tape position
// t.pos.
func (t *Tape) StoreByte(b byte) {
	t.tape[t.pos] = b
}
