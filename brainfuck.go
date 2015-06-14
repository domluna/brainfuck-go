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
)

var verbose = flag.Bool("verbose", false, "prints debug output")

// var usage = `
// `

var conf *config.Config

func usage() {
}

func main() {
	flag.Parse()

	if *verbose {
		conf = config.New(true)
	}

	if flag.NArg() > 0 {
		name := flag.Arg(1)

		var file io.Reader
		var err error

		file, err = os.Open(name)
		if err != nil {
			log.Fatalf("brainfuck: %s\n", err)
		}

		lexer := lex.New(name, conf, bufio.NewReader(file))
		parser := parse.New(name, conf, lexer)

		prog, err := parser.Parse()
		if err != nil {
			log.Fatalf("brainfuck: %s\n", err)
		}

		prog.Run()
	}
}