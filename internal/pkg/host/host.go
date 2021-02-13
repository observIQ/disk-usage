package host

import (
    "os"
    "net"
    "fmt"

    "github.com/pkg/errors"
)

// HostnameORAddr returns the hostname or primary ip address of the local system
func HostnameORAddr() (string, error) {
    h, err := getHostname()
    if err != nil {
        ip, err2 := getOutboundIP()
        if err2 != nil {
            return "", errors.Wrap(err, err.Error())
        }
        return ip, nil
    }
    return h, nil
}

func getHostname() (string, error) {
	h, err := os.Hostname()
    if err != nil {
        return "", err
    }
    if h == "" {
        return "", fmt.Errorf("could not get local hostname")
    }
    return h, nil
}

func getOutboundIP() (string, error) {
    // dial will not attempt to establish a connection, but it will return
    // the IP address that the system uses to talk to the outside world
    c, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        return "", err
    }
    defer c.Close()

    addr := c.LocalAddr().(*net.UDPAddr)
    ip := addr.IP.String()
    if ip == "" {
        return "", fmt.Errorf("could not get primary ip address")
    }
    return ip, nil
}
