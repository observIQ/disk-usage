package lockfile

import (
    "os"
    log "github.com/golang/glog"
    "github.com/pkg/errors"
)

// CreateLock takes a filepath and creates an empty
// file
func CreateLock(x string) error {
	f, err := os.Create(x)
    if err != nil {
		return errors.Wrap(err, "Failed to create lock file")
	}
    f.Close()
    return nil
}

// RemoveLock takes a filepath and removes the file
func RemoveLock(x string) error {
	if err := os.Remove(x); err != nil {
        return errors.Wrap(err, "Failed to remove lockfile")
	}

	log.Info("Lock removed")
	return nil
}

// Exists returns true if the lock file
// is present on the filesystem
func Exists(x string) bool {
    if _, err := os.Stat(x); err != nil {
        return false
    }
    return true
}
