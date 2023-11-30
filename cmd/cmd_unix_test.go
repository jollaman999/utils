// Run command for Linux & Unix like systems

//go:build !windows

package cmd

import "testing"

func Test_runCMD(t *testing.T) {
	_, err := RunCMD("ls -al")
	if err != nil {
		t.Fatal("Failed to run ls command!")
	}

	_, err = RunCMD("9 8 7 6 5 4 3 2 1")
	if err != nil {
		t.Log("Tried to run wrong command")
	}
}
