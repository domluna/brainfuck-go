// tool commands:
//	run
//	build
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/domluna/brainfuck-go/config"
	"github.com/domluna/brainfuck-go/lex"
	"github.com/domluna/brainfuck-go/parse"
	"github.com/domluna/brainfuck-go/program"
	"github.com/domluna/brainfuck-go/tape"
)

var conf = config.New(false)

// var usage = `
// `

// Hack around *bufio.Writer for os.Stdout.
// We call Flush after every write.
//
// Need this for interactive output and programs
// like rot13 which take input and return an output.
type byteWriterFlusher struct {
	w *bufio.Writer
}

func (bw byteWriterFlusher) WriteByte(b byte) error {
	err := bw.w.WriteByte(b)
	bw.w.Flush()
	return err
}

var usage = `Usage:

brainfuck-go <file>

Only takes 1 argument. All others are ignored.
`

func main() {
	flag.Parse()

	if flag.NArg() > 0 {
		name := flag.Arg(0)

		var file io.Reader
		var err error

		file, err = os.Open(name)
		if err != nil {
			log.Fatalf("brainfuck: reading program: %v", err)
		}

		lexer := lex.New(name, conf, bufio.NewReader(file))
		parser := parse.New(name, conf, lexer)

		prog := parser.Parse()

		// optimize program
		prog = program.Optimize(prog)

		// run program
		t := tape.New()
		in := bufio.NewReader(os.Stdin)
		out := byteWriterFlusher{bufio.NewWriter(os.Stdin)}
		for _, i := range prog {
			i.Eval(t, in, out)
		}

	} else {
		fmt.Print(usage)
	}
}
