package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"

	"github.com/domluna/brainfuck-go/config"
	"github.com/domluna/brainfuck-go/lex"
	"github.com/domluna/brainfuck-go/parse"
	"github.com/domluna/brainfuck-go/tape"
)

var verbose = flag.Bool("verbose", false, "prints debug output")

// var usage = `
// `

var conf *config.Config

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

func usage() {
}

func main() {
	flag.Parse()

	// if *verbose {
	// 	conf = config.New(true)
	// }
	// conf = config.New(true)

	if flag.NArg() > 0 {
		name := flag.Arg(0)

		var file io.Reader
		var err error

		file, err = os.Open(name)
		if err != nil {
			log.Fatalf("brainfuck: %s\n", err)
		}

		lexer := lex.New(name, conf, bufio.NewReader(file))
		parser := parse.New(name, conf, lexer)
		tape := tape.New()

		prog, err := parser.Parse()
		if err != nil {
			log.Fatalf("brainfuck: error during parsing %s\n", err)
		}

		in := bufio.NewReader(os.Stdin)
		out := byteWriterFlusher{bufio.NewWriter(os.Stdout)}

		prog.Run(tape, in, out)
	}
}
