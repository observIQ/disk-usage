package stdout

import (
	"fmt"
)

// Stdout type will send messages to standard out
type Stdout struct {
}

// Init will perform a no-op, required to satisfy the Alert interface
func (s Stdout) Init() error {
	return nil
}

// Send will send messages to standard out
func (s Stdout) Send(m string) error {
	fmt.Println("m")
	return nil
}
