// +build linux darwin

package main
import (
    "syscall"
    "fmt"
    "strconv"

    "github.com/bluemedorapublic/gopsutil/disk"
)


const lockpath string = "/tmp/suppress"


func getMountpoints() {
    devices, _ := disk.Partitions(true)
	for _, device := range devices {
        if checkFileSystem(device.Fstype) == true {
            drives = append(drives, device.Mountpoint)
        }
	}
}


func getUsage() {
    var (
        createAlert bool   = false
        createLock  bool   = false
        message     string = getHostname()
    )

    var stat syscall.Statfs_t
    fs  := syscall.Statfs_t{}

    for _, path := range drives {
        syscall.Statfs(path, &stat)
        err := syscall.Statfs(path, &fs)
        if err != nil {
            fmt.Println("Failed to read path:", path)

        } else {
            all  := int(fs.Blocks * uint64(fs.Bsize))
            free := int(fs.Bfree * uint64(fs.Bsize))
            used := int(all - free)
            percentage := int((float64(used) / float64(all)) * 100)

            if percentage > globalConfig.Threshold {
                fmt.Println("High disk usage:", path, strconv.Itoa(percentage) + "%")
                message = message + " high disk usage on drive " + path + " " + strconv.Itoa(percentage) + "% \n"
                createAlert = true
                createLock  = true

            } else {
                fmt.Println("Disk usage healthy:", path)

            }
        }
    }

    // If disk usage is healthy, and lock exists, clear it
    if createLock == false && lockExists(lockpath) {
        createAlert = true
    }

    if createAlert == true {
        if createLock == true {
            alert(message, true)
        } else {
            message = message + " disk usage cleared."
            removeLock(lockpath)  /* Remove the lock before alerting */
            alert(message, false)

        }
    }
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
    case "zfs":
        return true
    default:
        return false
    }
}
