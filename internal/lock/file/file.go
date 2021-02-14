package file

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
)

// File type manages a file at a given path
type File struct {
	path string
}

// New returns a new File type
func New(p string) File {
	return File{p}
}

// Lock will lock the file
func (f File) Lock() error {
	return lock(f.path)
}

// Unlock will unlock (delete) the file
func (f File) Unlock() error {
	return unlock(f.path)
}

// Exists will return true if the file exists
func (f File) Exists() bool {
	return exists(f.path)
}

// Path will return the file's path
func (f File) Path() string {
	return f.path
}

// Lock takes a filepath and creates an empty
// file
func lock(x string) error {
	f, err := os.Create(x)
	if err != nil {
		return errors.Wrap(err, "Failed to create lock file")
	}
	return f.Close()
}

// Unlock takes a filepath and removes the file
func unlock(x string) error {
	if err := os.Remove(x); err != nil {
		return errors.Wrap(err, "Failed to remove lockfile")
	}

	log.Info("Lock removed")
	return nil
}

// Exists returns true if the lock file
// is present on the filesystem
func exists(x string) bool {
	if _, err := os.Stat(x); err != nil {
		return false
	}
	return true
}
