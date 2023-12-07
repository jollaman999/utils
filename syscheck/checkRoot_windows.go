// Check root for Windows

//go:build windows

package syscheck

import (
	"errors"
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"syscall"
)

func runAsAdministrator() error {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		return err
	}

	return nil
}

func amAdmin() bool {
	return windows.GetCurrentProcessToken().IsElevated()
}

// CheckRoot : Run as Administrator in Windows with UAC
func CheckRoot() error {
	if !amAdmin() {
		err := runAsAdministrator()
		if err != nil {
			return errors.New("Failed to run as administrator. (" + err.Error() + ")")
		}
	}

	return nil
}
