package disk

import (
	"fmt"
	"encoding/json"

	"github.com/BlueMedoraPublic/disk-usage/internal/alert"
	"github.com/BlueMedoraPublic/disk-usage/internal/backend"

	log "github.com/sirupsen/logrus"
	"github.com/pkg/errors"
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

	// State interface
	State backend.State

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
	Healthy      bool   `json:"healthy"`
}

// State represents the state written to the state backend
type State struct {
	Alerted []string `json:"alerted"`
	Host  System   `json:"host"`
}

// Run will execute disk usage checks and alerts
func (c *Config) Run() error {
	if err := c.getDisks(); err != nil {
		return err
	}

	if err := c.checkUsage(); err != nil {
		return err
	}

	return c.handleState()
}

func (c *Config) checkUsage() error {
	log.Trace(fmt.Sprintf("Checking disk usage with threshold %d%%", c.Threshold))

	if err := c.getUsage(); err != nil {
		return err
	}

	for i, device := range c.Host.Devices {
		if device.UsagePercent > c.Threshold {
			c.Host.Devices[i].Healthy = false
			log.Trace(fmt.Sprintf("Device disk usage is unhealthy %s", device.Name))
			continue
		}
		c.Host.Devices[i].Healthy = true
		log.Trace(fmt.Sprintf("Device disk usage is healthy %s", device.Name))
	}
	return nil
}

func (c *Config) handleState() error {
	// Continue on failure, assume not alerted. Start fresh.
	prevState, err := c.ReadState()
	if err != nil {
		log.Error(errors.Wrap(err, "Starting fresh with new state"))
	}

	newState := State{
		Host: c.Host,
	}

	for _, current := range c.Host.Devices {

		if current.Healthy {
			m := fmt.Sprintf("device is healthy: %s", current.Name)
			log.Info(m)
			if prevState.alerted(current.Name) {
				if err := c.message(m, INFO); err != nil {
					// if alert fails, add device to new state as alerted
					log.Error(err)
					newState.Alerted = append(newState.Alerted, current.Name)
				}
				log.Info(fmt.Sprintf("Sent 'device is healthy' notification for device %s", current.Name))
			}
		}

		if ! current.Healthy {
			m := fmt.Sprintf("device is unhealthy: %s", current.Name)
			log.Warning(m)
			if ! prevState.alerted(current.Name) {
				if err := c.message(m, FATAL); err != nil {
					log.Error(err)
					// when alert fails, do not add device to state as alerted
					// in order to force the alert attempt next time disk-usage
					// is executed
					continue
				}
				log.Info(fmt.Sprintf("Sent 'device is unhealthy' alert for device %s", current.Name))
			}

			// add the device to the new state after alerting or skipping
			// due to already alerted
			newState.Alerted = append(newState.Alerted, current.Name)
		}
	}

 	if err := c.WriteState(newState); err != nil {
		return errors.Wrap(err, "Failed to write state")
	}
	return nil
}

func (c Config) ReadState() (State, error) {
	s := State{}
	b, err := c.State.Read()
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(b, &s)
	return s, err
}

func (c Config) WriteState(s State) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return c.State.Write(b)
}

func (s State) alerted(input string) bool {
	for _, i := range s.Alerted {
		if input == i {
			return true
		}
	}
	return false
}
