package syscheck

import (
	"errors"
	"os/exec"
)

// CheckMySQL : Check if we are ready to use mysql command
func CheckMySQL() error {
	cmd := exec.Command("mysql", "--help")
	err := cmd.Run()
	if err != nil {
		return errors.New("mysql command not found")
	}

	return nil

}
