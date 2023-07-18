package iputil

import (
	"os/exec"
)

// EnableIPForwardV4 : Enable IPv4 IP forward to use NAT with iptables.
func EnableIPForwardV4() error {
	cmd := exec.Command("echo", "1", ">", "/proc/sys/net/ipv4/ip_forward")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
