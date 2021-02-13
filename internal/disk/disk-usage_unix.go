// +build linux darwin freebsd

package disk

import (
	"strconv"
	"syscall"

	"github.com/bluemedorapublic/gopsutil/disk"
	log "github.com/golang/glog"
)

func (c *Config) getDisks() error {
	devices, err := disk.Partitions(true)
	if err != nil {
		return err
	}

	for _, device := range devices {
		if checkFileSystem(device.Fstype) == true {

			d := Device{
				Name: device.Device,
				MountPoint: device.Mountpoint,
				Type: device.Fstype,
			}
			c.Host.Devices = append(c.Host.Devices, d)

			c.Host.Drives = append(c.Host.Drives, device.Mountpoint)
		}
	}

	return nil
}

func (c Config) getUsage() error {
	var (
		createAlert bool   = false
		createLock  bool   = false
		message     string = c.Host.Name
	)

	var stat syscall.Statfs_t
	fs := syscall.Statfs_t{}

	for _, path := range c.Host.Drives {
		syscall.Statfs(path, &stat)
		err := syscall.Statfs(path, &fs)
		if err != nil {
			log.Info("Failed to read path:", path)

		} else {
			all := int(fs.Blocks * uint64(fs.Bsize))
			free := int(fs.Bfree * uint64(fs.Bsize))
			used := int(all - free)
			percentage := int((float64(used) / float64(all)) * 100)

			if percentage > c.Threshold {
				message = message + " high disk usage on drive " + path + " " + strconv.Itoa(percentage) + "% \n"
				log.Info(message)
				createAlert = true
				createLock = true

			} else {
				log.Info("Disk usage healthy:", path)
			}
		}
	}

	return c.handleLock(createLock, createAlert, message)
}

func checkFileSystem(fs string) bool {
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
