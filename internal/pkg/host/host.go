package host

import (
	"fmt"
	"net"
)

// PrimaryAddress address will return the primary ip address of the system
func PrimaryAddress() (string, error) {
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
