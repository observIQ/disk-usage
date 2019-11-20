// +build windows

package main
import (
    "strconv"

    log "github.com/golang/glog"
    "github.com/bluemedorapublic/gopsutil/disk"
)


const lockpath string = "C:\\suppress.txt"


// Call the Partitions function to get an array all drevices (local disk, remote, usb, cdrom)
// Only append valid drvies to the drives array (Only local disks)
func getMountpoints() error {
	devices, err := disk.Partitions(true)
    if err != nil {
        return err
    }

	for _, device := range devices {
		if validDrive(int(device.Typeret)) == true {
			drives = append(drives, string(device.Mountpoint))
		}
	}

    return nil
}


// Kick off an alert for each drive that has a high consumption
func getUsage() error {
    var (
        createAlert bool   = false
        createLock  bool   = false
        message     string = globalConfig.Hostname
    )

    for _, drive := range drives {

		fs, _ := disk.Usage(drive + "\\")
        log.Info(fs.Path, int(fs.UsedPercent), "%")
		usedSpace := strconv.Itoa(int(fs.UsedPercent)) + "%"

		if int(fs.UsedPercent) > globalConfig.Threshold {
            message = message + " high disk usage on drive " + drive + " " + usedSpace
            log.Info(message)
            createAlert = true
            createLock = true

        } else {
            log.Info("Disk usage healthy")
        }
    }

    return handleLock(createLock, createAlert, message)
}


func getDevType(driveType uintptr) string {
	switch driveType {
	case 2:
		return "Removable"
	case 3:
		return "Local"
	case 4:
		return "Network"
	case 5:
		return "CDROM"
	default:
		return "Unknown"
	}
}


// Local drives are the only drives that should be considered for alerting
func validDrive(driveType int) bool {
	if driveType == 3 {
		return true
	} else {
		return false
	}
}

func lockPath() string {
	return lockpath
}
