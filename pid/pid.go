package pid

import (
	"github.com/jollaman999/utils/fileutil"
	"os"
	"strconv"
	"syscall"
)

// IsModuleRunning : Check if module is running
func IsModuleRunning(moduleName string) (running bool, pid int, err error) {
	var PIDFile = "/run/" + moduleName + ".pid"

	if _, err := os.Stat(PIDFile); os.IsNotExist(err) {
		return false, 0, nil
	}

	pidStr, err := os.ReadFile(PIDFile)
	if err != nil {
		return false, 0, err
	}

	PID, _ := strconv.Atoi(string(pidStr))

	proc, err := os.FindProcess(PID)
	if err != nil {
		return false, 0, err
	}
	err = proc.Signal(syscall.Signal(0))
	if err == nil {
		return true, PID, nil
	}

	return false, 0, nil
}

// WriteHarpPID : Write harp PID to "/var/run/harp.pid"
func WriteHarpPID(moduleName string) error {
	var PIDFile = "/run/" + moduleName + ".pid"

	pid := os.Getpid()

	err := fileutil.CreateDirIfNotExist("/run")
	if err != nil {
		return err
	}

	err = fileutil.WriteFile(PIDFile, strconv.Itoa(pid))
	if err != nil {
		return err
	}

	return nil
}

// DeleteHarpPID : Delete the harp PID file
func DeleteHarpPID(moduleName string) error {
	var PIDFile = "/run/" + moduleName + ".pid"

	err := fileutil.DeleteFile(PIDFile)
	if err != nil {
		return err
	}

	return nil
}
