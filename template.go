// template used to generate a Go program for compilation to binary.
package main

import "text/template"

var generatedTmpl = template.Must(template.New("generated").Parse(`
// generated by brainfuck-go; DO NOT EDIT

package main

import (
	"bufio"
	"os"

	"github.com/domluna/brainfuck-go/tape"
)

func main() {
	tape := tape.New()
	in := bufio.NewReader(os.Stdin)
	out := byteWriterFlusher{bufio.NewWriter(os.Stdout)}
	{{range _, $inst := .Insts}}
		$inst.Eval(tape, in, out)
	{{end}}
}
`))
