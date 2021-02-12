package null

// Null type performs no operations
type Null struct {
}

// Lock will perform a no-op when called
func (n Null) Lock() error {
	return nil
}

// Unlock will perform a no-op when called
func (n Null) Unlock() error {
	return nil
}

// Exists will perform a no-op when called
func (n Null) Exists() bool {
	return false
}

// Path will perform a no-op when called
func (n Null) Path() string {
	return ""
}
