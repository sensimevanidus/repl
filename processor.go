package repl

// Processor defines the interface that is to be implemented as a concrete processor that will act as the "eval" part
// of the REPL.
type Processor interface {
	// Process method is called in the "loop" part of the REPL. The input is passed, and the output is expected.
	// TODO: We may think to use channels for the communication between the processor and the loop. That way, we could
	// have a bi-directional communication.
	Process(input []byte) (string, error)
}
