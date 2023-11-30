// Check root for Linux & Unix like systems

//go:build !windows

package syscheck

import "testing"

func Test_Root(t *testing.T) {
	err := CheckRoot()
	if err != nil {
		t.Fatal("Failed to get root permission!")
	}
}
