# Brainfuck-go

Brainfuck compiler in Go - well not quite. A compiler outputs binaries. It's in
between an interpreter and a compiler.

### Intro

This is not the smallest possible implementation of a Brainfuck interpreter by any means. Nor it is the fastest. Those aren't the goals.

The goal from this was to learn a bit about how interpreters and compilers work. Using Brainfuck looked to be the simplest way to achieve this. 

Steps of this process are kept in separate packages and serve as a decent blueprint for future, more complicated work.

Resources:

[Lexical Scanning in Go](https://www.youtube.com/watch?v=HxaD_trXwRE)

[Ivy, APL Interpreter](https://github.com/robpike/ivy)

### How it works

1. lexer on *.b file
2. parses tokens from lexer
3. create list(AST in this context) of instructions from (2) 
4. optimize instructions
5. evalute instructions
6. output result

Steps 1-3 run concurrently.

### Usage

Get it.

```sh
$ go get github.com/domluna/brainfuck-go
```

Use it.

```sh
$ brainfuck-go helloworld.b
// Hello World!
```

### Optimizing the instructions

Running `examples/mandelbrot.b` timings.

```sh
    without optimization -> 194.10 seconds
    with    optimization -> 52.62  seconds
```

That's ~400% speed up. So it's fair to say the optimizations do a fair amount.


