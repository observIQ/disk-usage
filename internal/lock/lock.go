package lock

import (
	"fmt"

	"github.com/BlueMedoraPublic/disk-usage/internal/lock/file"
	"github.com/BlueMedoraPublic/disk-usage/internal/lock/null"
)

// Lock interface manages a locking backend to prevent alert storms
type Lock interface {
	Lock() error
	Unlock() error
	Exists() bool
	Path() string
}

// File will return a new File type as a Lock interface
func File(path string) (Lock, error) {
	if path == "" {
		return nil, fmt.Errorf("file lock path not set")
	}
	return file.New(path), nil
}

// Null will return a new Null type as a Lock interface
func Null() Lock {
	return null.Null{}
}
