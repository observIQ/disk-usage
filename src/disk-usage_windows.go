// +build windows

package main
import (
    "fmt"
    "strconv"

    "github.com/bluemedorapublic/gopsutil/disk"
)


const lockpath string = "C:\\suppress.txt"


// Call the Partitions function to get an array all drevices (local disk, remote, usb, cdrom)
// Only append valid drvies to the drives array (Only local disks)
func getMountpoints() {
	devices, _ := disk.Partitions(true)
	for _, device := range devices {
		if validDrive(int(device.Typeret)) == true {
			drives = append(drives, string(device.Mountpoint))
		}
	}
}


// Kick off an alert for each drive that has a high consumption
func getUsage() {
    for _, drive := range drives {

		fs, _ := disk.Usage(drive + "\\")
		fmt.Println(fs.Path, int(fs.UsedPercent), "%")
		usedSpace := strconv.Itoa(int(fs.UsedPercent)) + "%"

		if int(fs.UsedPercent) > globalConfig.Threshold {
            message := getHostname() + " high disk usage on drive " + drive + " " + usedSpace
			alert(message, true)

        } else {
            fmt.Println("Disk usage healthy")
            if lockExists(lockpath) {
                removeLock(lockpath)
                message := getHostname() + " disk usage cleared: " + drive + " " + usedSpace
                alert(message, false)
            }
        }
	}
}


func getDevType(typeret uintptr) string {
	switch typeret {
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
func validDrive(typeret int) bool {
	if typeret == 3 {
		return true
	} else {
		return false
	}
}
