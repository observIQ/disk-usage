package disk

import (
    "github.com/BlueMedoraPublic/disk-usage/internal/alert"
    "github.com/BlueMedoraPublic/disk-usage/internal/lockfile"

    log "github.com/golang/glog"
)

type Config struct {
	Threshold int
	Hostname string

	// alert interface
	Alert alert.Alert

    drives []string
}

func (c *Config) Run() error {
    if err := c.getMountpoints(); err != nil {
        return err
    }
    return c.getUsage()
}

func (c Config) handleLock(createLock, createAlert bool, message string) error {
	// If disk usage is healthy, and lock exists, clear it
	// by removing the lock
	if createLock == false && lockfile.Exists(lockPath()) {
		if err := c.Alert.Send(message + " disk usage is healthy"); err != nil {
			return err
		}
		return lockfile.RemoveLock(lockPath())
	}

	// If disk usage is not healthy and lockfile does not exist,
	// fire off an alert
	if createLock && !lockfile.Exists(lockPath()) {
		if err := c.Alert.Send(message); err != nil {
			return err
		}
		return lockfile.CreateLock(lockPath())
	}

	if createLock == true && lockfile.Exists(lockPath()) {
		log.Info("Lock exists, skipping alert.")
		return nil
	}

	return nil
}
