// Run command for Linux & Unix like systems

//go:build !windows

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

func RunSh(commandString string) (output string, err error) {
	return runCMD("sh", "-c", commandString)
}

func RunBash(commandString string) (output string, err error) {
	return runCMD("bash", "-c", commandString)
}

func RunZsh(commandString string) (output string, err error) {
	return runCMD("zsh", "-c", commandString)
}

func RunCMD(commandString string) (output string, err error) {
	return RunSh(commandString)
}
