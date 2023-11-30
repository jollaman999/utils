// Run command for Windows

//go:build windows

package cmd

import (
	"errors"
	"os/exec"
)

func runCMD(name string, arg ...string) (output string, err error) {
	cmd := exec.Command(name, arg...)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		if len(outputBytes) == 0 {
			outputBytes = []byte(err.Error())
		}
		return string(outputBytes), errors.New(string(outputBytes))
	}

	return string(outputBytes), nil
}

func RunPowerShell(commandString string) (output string, err error) {
	return runCMD("powershell.exe", "-Command", commandString)
}

func RunCMD(commandString string) (output string, err error) {
	return runCMD("cmd.exe", "/c", commandString)
}
