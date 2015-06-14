# Brainfuck-go

Brainfuck interpreter/compiler in Go.

How to deal with output?

WriteByte will write the byte to some output.
Should it be

WriteByte()
or
WriteByte(w io.ByteWriter)
?

Options:

1. Add another argument to Eval(), like the Program. We could then add a field to the Program
type and use WriteByte().

2. For WriteByte(io.ByteWriter) we make a field in the instruction. 

3. Add a field to the type that implements Tape, and use WriteByte(), if we don't care
this could also just print to stdout.

Which one?

1. Problem here we add an extra argument but InstWriteByte would be the only instruction
to make use of this. Not a good solution.

2. Adding a field to the instruction might seem reasonable. But, we have to make sure
the field is a pointer. Even still, we'd have hundreds or thousands of instructions that
point to the same io.ByteWriter.

3. Best option. We remove the output completely from the program package and let whoever is
implementing the interface worry about it.

