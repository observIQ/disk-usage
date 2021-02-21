package null

// Null type performs no operations
type Null struct {
}

// Write will perform a no-op when called
func (n Null) Write(state []byte) error {
	return nil
}

// Read will perform a no-op when called
func (n Null) Read() ([]byte, error) {
	return nil, nil
}

// Path will perform a no-op when called
func (n Null) Path() string {
	return ""
}
