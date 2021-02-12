package disk

import (
    "github.com/BlueMedoraPublic/disk-usage/internal/alert"
    "github.com/BlueMedoraPublic/disk-usage/internal/lock"

    log "github.com/golang/glog"
)

type Config struct {
	Threshold int
	Hostname string

	// alert interface
	Alert alert.Alert

    // lock interface
    Lock lock.Lock

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
	if !createLock && c.Lock.Exists() {
		if err := c.Alert.Send(message + " disk usage is healthy"); err != nil {
			return err
		}
		return c.Lock.Unlock()
	}

	// If disk usage is not healthy and lock does not exist,
	// fire off an alert
	if createLock && !c.Lock.Exists() {
		if err := c.Alert.Send(message); err != nil {
			return err
		}
		return c.Lock.Lock()
	}

	if createLock == true && c.Lock.Exists() {
		log.Info("Lock exists, skipping alert.")
		return nil
	}

	return nil
}
