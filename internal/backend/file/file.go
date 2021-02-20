package file

import (
	"os"

	"github.com/pkg/errors"
)

// File type manages a file at a given path
type File struct {
	path string
}

// New returns a new File type
func New(p string) File {
	return File{p}
}

// Write will lock the file
func (f File) Write(state []byte) error {
	if err := os.WriteFile(f.path, state, 0644); err != nil {
		return errors.Wrap(err, "Failed to write state file")
	}
	return nil
}

// Read will read the file
func (f File) Read() ([]byte, error) {
	b, err := os.ReadFile(f.path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read state file")
	}
	return b, nil
}

// Path will return the file's path
func (f File) Path() string {
	return f.path
}
