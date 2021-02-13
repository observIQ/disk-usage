package disk

import (
	"encoding/json"

	"github.com/BlueMedoraPublic/disk-usage/internal/alert"
	"github.com/BlueMedoraPublic/disk-usage/internal/lock"

	log "github.com/golang/glog"
)

const INFO = "info"
const FATAL = "fatal"

// Config type represents the configuration for
// disk-usage checking and alerting
type Config struct {
	Threshold int

	// alert interface
	Alert alert.Alert

	// lock interface
	Lock lock.Lock

	Host System
}

type System struct {
	Name    string   `json:"name"`
	Drives  []string `json:"drives"`
	Devices []Device `json:"devices"`
}

type Message struct {
	Host     System `json:"host"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
}

type Device struct {
	Name       string `json:"name"`
	MountPoint string `json:"mountpoint"`
	Type       string `json:"type"`
}

// Run will execute disk usage checks and alerts
func (c *Config) Run() error {
	if err := c.getDisks(); err != nil {
		return err
	}
	return c.getUsage()
}

func (c Config) handleLock(createLock, createAlert bool, message string) error {
	// If disk usage is healthy, and lock exists, clear it
	// by removing the lock
	if !createLock && c.Lock.Exists() {
		m := message + " disk usage is healthy"
		if err := c.message(m, INFO); err != nil {
			return err
		}
		return c.Lock.Unlock()
	}

	// If disk usage is not healthy and lock does not exist,
	// fire off an alert
	if createLock && !c.Lock.Exists() {
		if err := c.message(message, FATAL); err != nil {
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

func (c Config) message(msg, sev string) error {
	m := Message{
		Host: c.Host,
		Message: msg,
		Severity: sev,
	}

	b, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return err
	}

	return c.Alert.Send(string(b))
}
