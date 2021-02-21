package backend

import (
	"fmt"

	"github.com/BlueMedoraPublic/disk-usage/internal/backend/file"
	"github.com/BlueMedoraPublic/disk-usage/internal/backend/null"
)

// State interface manages writing and reading state
type State interface {
	Write(state []byte) error
	Read() ([]byte, error)
	Path() string
}

// File will return a new File type as a State interface
func File(path string) (State, error) {
	if path == "" {
		return nil, fmt.Errorf("file state path not set")
	}
	return file.New(path), nil
}

// Null will return a new Null type as a state interface
func Null() State {
	return null.Null{}
}
