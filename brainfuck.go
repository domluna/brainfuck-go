// tool commands:
//	run
//	build
//
package main

import (
	"bufio"
	"bytes"
	"flag"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/domluna/brainfuck-go/config"
	"github.com/domluna/brainfuck-go/lex"
	"github.com/domluna/brainfuck-go/parse"
	"github.com/domluna/brainfuck-go/program"
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
			log.Fatalf("brainfuck: reading program: %v", err)
		}

		lexer := lex.New(name, conf, bufio.NewReader(file))
		parser := parse.New(name, conf, lexer)

		prog, err := parser.Parse()
		if err != nil {
			log.Fatalf("brainfuck: parsing: %v", err)
		}

		in := bufio.NewReader(os.Stdin)
		out := byteWriterFlusher{bufio.NewWriter(os.Stdout)}
		insts := prog.Insts

		var compiled = struct {
			Insts []program.Instruction
		}{
			Insts: insts,
		}

		var buf bytes.Buffer
		if err := generatedTmpl.Execute(&buf, compiled); err != nil {
			log.Fatalf("brainfuck: generating code: %v", err)
		}

		src, err := format.Source(buf.Bytes())
		if err != nil {
			log.Printf("warning: invalid Go generated: %s", err)
			log.Printf("warning: compile code to see error")
			src = buf.Bytes()
		}

		dir := "./code"
		outFile := strings.ToLower(name + "_brainfuck.go")
		outPath := filepath.Join(dir, outFile)
		if err := ioutil.WriteFile(outPath, src, 0644); err != nil {
			log.Fatalf("brainfuck: writing output file: %s", err)
		}

	}
}
