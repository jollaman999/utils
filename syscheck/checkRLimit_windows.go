// Increasing RLimit is not supported on Windows

//go:build windows

package syscheck

import "errors"

// IncreaseRLimitToMax : Increase rLimit to system max
func IncreaseRLimitToMax() error {
	return errors.New("Increasing RLimit is not supported on Windows")
}
