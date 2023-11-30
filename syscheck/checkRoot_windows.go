// Check root for Windows

//go:build windows

package syscheck

import (
	"errors"
)

// CheckRoot : Check root permission
func CheckRoot() error {
	return errors.New("checking root is not supported on Windows")
}
