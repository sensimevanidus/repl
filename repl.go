package repl

import (
	"bufio"
	"fmt"
	"io"
)

const (
	prompt = ">> "
)

// Run is the actual "loop" part of the REPL. It basically reads the user input (R), passes that to the processor's
// `Process` method which evaluates it (E) and returns the output. The output is printed to the given writer (P) and
// the loop goes back to the waiting user input operation (L).
func Run(in io.Reader, out io.Writer, inputProcessor Processor) error {
	scanner := bufio.NewScanner(in)

	for {
		// prompt for input
		out.Write([]byte(prompt))

		// wait for input
		if ok := scanner.Scan(); !ok {
			// scan stopped (end of the input or an error occurred)
			return scanner.Err()
		}

		// process the input, and fetch the output
		if output, err := inputProcessor.Process(scanner.Bytes()); err != nil {
			out.Write([]byte(fmt.Sprintf("[ERROR] %+v\n", err)))
		} else {
			out.Write([]byte(fmt.Sprintf("%v\n", output)))
		}
	}
}
