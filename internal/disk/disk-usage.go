package disk

import (
	"fmt"


	"github.com/BlueMedoraPublic/disk-usage/internal/alert"
	"github.com/BlueMedoraPublic/disk-usage/internal/lock"

	log "github.com/golang/glog"
)

// INFO represents info severity
const INFO = "info"
// FATAL represents fatal severity
const FATAL = "fatal"

// Config type represents the configuration for
// disk-usage checking and alerting
type Config struct {
	// Threshold is the percentage disk usage required to driver an alert
	Threshold int

	// alert interface
	Alert alert.Alert

	// lock interface
	Lock lock.Lock

	// Host is the system being managed by this config
	Host System
}

// System represents the local system
type System struct {
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Devices []Device `json:"devices"`
}

// Device represents a device attached to the system
type Device struct {
	Name         string `json:"name"`
	MountPoint   string `json:"mountpoint"`
	Type         string `json:"type"`
	UsagePercent int    `json:"usage_percent"`
}

// Run will execute disk usage checks and alerts
func (c *Config) Run() error {
	if err := c.getDisks(); err != nil {
		return err
	}
	return c.checkUsage()
}

func (c Config) checkUsage() error {
	if err := c.getUsage(); err != nil {
		return err
	}

	highUsage := []string{}
	for _, device := range c.Host.Devices {
		if device.UsagePercent > c.Threshold {
			highUsage = append(highUsage, device.Name)
		}
	}

	if len(highUsage) > 0 {
		msg := fmt.Sprintf("devices have high usage: %s", highUsage)
		return c.handleLock(true, msg)
	}
	return c.handleLock(false, "disk usage healthy")
}

func (c Config) handleLock(unhealthy bool, message string) error {
	// If disk usage is healthy, and lock exists, clear it
	// by removing the lock
	if !unhealthy && c.Lock.Exists() {
		m := message + " disk usage is healthy"
		if err := c.message(m, INFO); err != nil {
			return err
		}
		return c.Lock.Unlock()
	}

	// If disk usage is not healthy and lock does not exist,
	// fire off an alert
	if unhealthy && !c.Lock.Exists() {
		if err := c.message(message, FATAL); err != nil {
			return err
		}
		return c.Lock.Lock()
	}

	if unhealthy == true && c.Lock.Exists() {
		log.Info("Lock exists, skipping alert.")
		return nil
	}

	return nil
}
