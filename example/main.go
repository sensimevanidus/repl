package main

import (
	"fmt"
	"os"

	"github.com/sensimevanidus/repl"
)

// InputPrinter is an example implementation for the repl.Processor interface.
type InputPrinter struct{}

// Process prints the input to the stdout as-is.
func (inputPrinter InputPrinter) Process(input []byte) (string, error) {
	return fmt.Sprintf("%v", string(input)), nil
}

func main() {
	if err := repl.Run(os.Stdin, os.Stdout, InputPrinter{}); err != nil {
		fmt.Errorf("%v\n", err.Error())
		os.Exit(1)
	}
}
