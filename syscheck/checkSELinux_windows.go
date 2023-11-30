// Check SELinux for Linux

//go:build linux

package syscheck

func CheckSELinuxEnforced() bool {
	return false
}
