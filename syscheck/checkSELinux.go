package syscheck

import (
	"github.com/opencontainers/selinux/go-selinux"
)

func CheckSELinuxEnforced() bool {
	return selinux.GetEnabled()
}
