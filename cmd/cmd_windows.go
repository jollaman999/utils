// Run command for Windows

//go:build windows

package cmd

import (
	"errors"
	"os/exec"
)

func runCMD(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if len(output) == 0 {
			output = []byte(err.Error())
		}
		return errors.New(string(output))
	}

	return nil
}

func RunPowerShell(commandString string) (output string, err error) {
	return runCMD("powershell.exe", "-Command", commandString)
}

func RunCMD(commandString string) (output string, err error) {
	return runCMD("cmd.exe", "/c", commandString)
}
