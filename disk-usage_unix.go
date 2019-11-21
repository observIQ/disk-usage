// +build linux darwin freebsd

package main

import (
	"strconv"
	"syscall"

	log "github.com/golang/glog"
	"github.com/bluemedorapublic/gopsutil/disk"
)

const lockpath string = "/tmp/suppress"

func getMountpoints() error {
	devices, err := disk.Partitions(true)
	if err != nil {
		return err
	}

	for _, device := range devices {
		if checkFileSystem(device.Fstype) == true {
			drives = append(drives, device.Mountpoint)
		}
	}

	return nil
}

func getUsage() error {
	var (
		createAlert bool   = false
		createLock  bool   = false
		message     string = globalConfig.Hostname
	)

	var stat syscall.Statfs_t
	fs := syscall.Statfs_t{}

	for _, path := range drives {
		syscall.Statfs(path, &stat)
		err := syscall.Statfs(path, &fs)
		if err != nil {
			log.Info("Failed to read path:", path)

		} else {
			all := int(fs.Blocks * uint64(fs.Bsize))
			free := int(fs.Bfree * uint64(fs.Bsize))
			used := int(all - free)
			percentage := int((float64(used) / float64(all)) * 100)

			if percentage > globalConfig.Threshold {
				message = message + " high disk usage on drive " + path + " " + strconv.Itoa(percentage) + "% \n"
				log.Info(message)
				createAlert = true
				createLock = true

			} else {
				log.Info("Disk usage healthy:", path)
			}
		}
	}

	return handleLock(createLock, createAlert, message)
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

func lockPath() string {
	return lockpath
}
