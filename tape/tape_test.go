package tape_test

import (
	"bytes"
	"testing"

	"github.com/domluna/brainfuck-go/tape"
)

func TestBasics(t *testing.T) {
	var out bytes.Buffer
	var b byte
	tt := tape.New(&out)

	b = tt.GetByte()
	if b != byte(0) {
		t.Fatalf("GetByte(): expected 0, got %v", b)
	}

	tt.MoveHead(1)
	if tt.GetHead() != 1 {
		t.Fatalf("MoveHead(1) from 0: expected 1, got %v", tt.GetHead())
	}

	tt.AddToByte(-1)
	t.Log(tt)
	b = tt.GetByte()
	if b != byte(255) {
		t.Fatalf("AddToByte(-1) from 0: expected 255, got %v", b)
	}

	tt.AddToByte(1)
	t.Log(tt)
	b = tt.GetByte()
	if b != byte(0) {
		t.Fatalf("AddToByte(1) from 255: expected 0, got %v", b)
	}

	tt.MoveHead(1)
	t.Log(tt)
	if tt.GetHead() != 2 {
		t.Fatalf("MoveHead(1) from 1: expected 2, got %v", tt.GetHead())
	}

	tt.SetByte(byte(100))
	t.Log("Expecting tape of [0 0 100]")

	b = tt.GetByte()
	if b != byte(100) {
		t.Fatalf("StoreByte(100): expected 100, got %v", b)
	}

	tt.MoveHead(-2)
	t.Logf("tape head at %d", tt.GetHead())
	b = tt.GetByte()
	if b != byte(0) {
		t.Fatalf("MoveHead(-2) from 2: expected 0, got %v", b)
	}

}

// test overflows and underflows
func TestFlows(t *testing.T) {
	var out bytes.Buffer
	var b byte
	tt := tape.New(&out)

	// testing panics
	tt.MoveHead(10)
	t.Log(tt)
	tt.GetHead()

	tt.MoveHead(-5)
	t.Log(tt)

	tt.MoveHead(6)
	t.Log(tt)
	tt.GetHead()

	tt.AddToByte(1000) // 232
	b = tt.GetByte()
	if b != 232 {
		t.Fatalf("AddToByte(1000) from 0: expected 232, got %v", b)
	}

	tt.AddToByte(-1000) // 0
	b = tt.GetByte()
	if b != 0 {
		t.Fatalf("AddToByte(-1000) from 1000: expected 0, got %v", b)
	}

	tt.AddToByte(1000)
	tt.AddToByte(-2000)
	b = tt.GetByte()
	if b != 24 {
		t.Fatalf("AddToByte(-2000) from 1000: expected 24, got %v", b)
	}
}
