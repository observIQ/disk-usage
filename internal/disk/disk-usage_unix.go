// +build linux darwin freebsd

package disk

import (
	"fmt"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/shirou/gopsutil/disk"
)

func (c *Config) getDisks() error {
	devices, err := disk.Partitions(true)
	if err != nil {
		return err
	}

	for _, device := range devices {
		if validDrive(device.Fstype) == true {
			d := Device{
				Name:       device.Device,
				MountPoint: device.Mountpoint,
				Type:       device.Fstype,
			}
			c.Host.Devices = append(c.Host.Devices, d)
		}
	}

	return nil
}

func (c *Config) getUsage() error {
	var stat syscall.Statfs_t
	fs := syscall.Statfs_t{}

	for i, device := range c.Host.Devices {
		path := device.MountPoint
		syscall.Statfs(path, &stat)
		err := syscall.Statfs(path, &fs)
		if err != nil {
			log.Error(fmt.Sprintf("failed to read path %s: %s", path, err.Error()))
			continue
		}

		all := int(fs.Blocks * uint64(fs.Bsize))
		free := int(fs.Bfree * uint64(fs.Bsize))
		used := int(all - free)
		percentage := int((float64(used) / float64(all)) * 100)
		c.Host.Devices[i].UsagePercent = percentage

		log.Trace(fmt.Sprintf("disk %s usage %d", device.Name, percentage))
	}
	return nil
}

func validDrive(fs string) bool {
	switch fs {
	case "xfs":
		return true
	case "ext4":
		return true
	case "ext3":
		return true
	case "ext2":
		return true
	case "ext":
		return true
	case "ufs":
		return true
	case "zfs":
		return true
	default:
		return false
	}
}
