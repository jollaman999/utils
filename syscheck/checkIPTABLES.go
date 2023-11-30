package syscheck

import (
	"errors"
	"os/exec"
)

// CheckIPTABLES : Check if we are ready to use iptables
func CheckIPTABLES() error {
	cmd := exec.Command("iptables", "--help")
	err := cmd.Run()
	if err != nil {
		return errors.New("iptables command not found")
	}

	return nil

}
