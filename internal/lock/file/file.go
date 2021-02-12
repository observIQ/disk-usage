package file

import (
    "os"
    log "github.com/golang/glog"
    "github.com/pkg/errors"
)

type File struct {
    path string
}

func New(p string) File {
    return File{p}
}

func (f File) Lock() error {
    return lock(f.path)
}

func (f File) Unlock() error {
    return unlock(f.path)
}

func (f File) Exists() bool {
    return exists(f.path)
}

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
