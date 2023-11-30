// Check SELinux for Linux

//go:build linux

package syscheck

import (
	"github.com/opencontainers/selinux/go-selinux"
)

func CheckSELinuxEnforced() bool {
	return selinux.GetEnabled()
}
