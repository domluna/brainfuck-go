package main_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/domluna/brainfuck-go/config"
	"github.com/domluna/brainfuck-go/lex"
	"github.com/domluna/brainfuck-go/parse"
	"github.com/domluna/brainfuck-go/program"
	"github.com/domluna/brainfuck-go/tape"
)

const helloWorldProg = `
++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.
`

const rot13Prog = `
-,+[                         Read first character and start outer character reading loop
    -[                       Skip forward if character is 0
        >>++++[>++++++++<-]  Set up divisor (32) for division loop
                               (MEMORY LAYOUT: dividend copy remainder divisor quotient zero zero)
        <+<-[                Set up dividend (x minus 1) and enter division loop
            >+>+>-[>>>]      Increase copy and remainder / reduce divisor / Normal case: skip forward
            <[[>+<-]>>+>]    Special case: move remainder back to divisor and increase quotient
            <<<<<-           Decrement dividend
        ]                    End division loop
    ]>>>[-]+                 End skip loop; zero former divisor and reuse space for a flag
    >--[-[<->+++[-]]]<[         Zero that flag unless quotient was 2 or 3; zero quotient; check flag
        ++++++++++++<[       If flag then set up divisor (13) for second division loop
                               (MEMORY LAYOUT: zero copy dividend divisor remainder quotient zero zero)
            >-[>+>>]         Reduce divisor; Normal case: increase remainder
            >[+[<+>-]>+>>]   Special case: increase remainder / move it back to divisor / increase quotient
            <<<<<-           Decrease dividend
        ]                    End division loop
        >>[<+>-]             Add remainder back to divisor to get a useful 13
        >[                   Skip forward if quotient was 0
            -[               Decrement quotient and skip forward if quotient was 1
                -<<[-]>>     Zero quotient and divisor if quotient was 2
            ]<<[<<->>-]>>    Zero divisor and subtract 13 from copy if quotient was 1
        ]<<[<<+>>-]          Zero divisor and add 13 to copy if quotient was 0
    ]                        End outer skip loop (jump to here if ((character minus 1)/32) was not 2 or 3)
    <[-]                     Clear remainder from first division if second division was skipped
    <.[-]                    Output ROT13ed character from copy and clear it
    <-,+                     Read next character
]                            End character reading loop
`

// only need Debug=true if we screw up and need to see what's going on
var conf = config.New(false)

func Test_HelloWorld(t *testing.T) {

	lexer := lex.New("hello_world.b", conf, strings.NewReader(helloWorldProg))
	parser := parse.New("hello_world.b", conf, lexer)

	prog := parser.Parse()

	var result string
	var tp *tape.Tape
	var in *strings.Reader
	var out bytes.Buffer

	tp = tape.New()
	in = strings.NewReader("")

	expect := "Hello World!\n"

	for _, i := range prog {
		i.Eval(tp, in, &out)
	}

	result = out.String()
	if result != expect {
		t.Errorf("normal program: expected %s, got %s", expect, result)
	}

	// reset
	tp = tape.New()
	in = strings.NewReader("")
	out.Reset()

	prog = program.Optimize(prog)
	for _, i := range prog {
		i.Eval(tp, in, &out)
	}
	result = out.String()
	if result != expect {
		t.Errorf("optimized program: expected %s, got %s", expect, result)
	}

}

func Test_Rot13(t *testing.T) {
	lexer := lex.New("rot13.b", conf, strings.NewReader(rot13Prog))
	parser := parse.New("rot13.b", conf, lexer)

	prog := parser.Parse()

	var result string
	var tp *tape.Tape
	var in *strings.Reader
	var out bytes.Buffer

	expect := "V'z gur ongzna!"

	tp = tape.New()
	in = strings.NewReader("I'm the batman!")

	for _, i := range prog {
		i.Eval(tp, in, &out)
	}

	result = out.String()
	if result != expect {
		t.Errorf("normal program: expected %s, got %s", expect, result)
	}

	// reset
	tp = tape.New()
	in = strings.NewReader("I'm the batman!")
	out.Reset()

	prog = program.Optimize(prog)
	for _, i := range prog {
		i.Eval(tp, in, &out)
	}

	result = out.String()
	if result != expect {
		t.Errorf("optimized program: expected %s, got %s", expect, result)
	}

}
