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

func RunSh(commandString string) error {
	return runCMD("sh", "-c", commandString)
}

func RunBash(commandString string) error {
	return runCMD("bash", "-c", commandString)
}

func RunZsh(commandString string) error {
	return runCMD("zsh", "-c", commandString)
}

func RunCMD(commandString string) error {
	return RunSh(commandString)
}
