// Load kernel modules for Linux & Unix like systems

//go:build !windows

package modprobe

import "pault.ag/go/modprobe"

func Load(moduleName string) error {
	return modprobe.Load(moduleName, "")
}

func Unload(moduleName string) error {
	return modprobe.Remove(moduleName)
}
