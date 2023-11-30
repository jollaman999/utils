// PID tool for Linux & Unix like systems

//go:build !windows

package pid

import (
	"github.com/jollaman999/utils/fileutil"
	"os"
	"strconv"
	"syscall"
)

// CheckRunningWithPIDFile : Check if process is running with PID file.
func CheckRunningWithPIDFile(PIDFile string) (running bool, pid int, err error) {
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

// WritePIDFile : Write PID file of currently running module.
func WritePIDFile(PIDFile string) error {
	pid := os.Getpid()

	err := fileutil.WriteFile(PIDFile, strconv.Itoa(pid))
	if err != nil {
		return err
	}

	return nil
}
