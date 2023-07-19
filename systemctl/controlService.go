package systemctl

import (
	"github.com/jollaman999/utils/cmd"
	"github.com/jollaman999/utils/syscheck"
)

// RestartService : Restart the service with systemctl command.
func RestartService(serviceName string) error {
	err := syscheck.CheckRoot()
	if err != nil {
		return err
	}

	err = cmd.RunCMD("systemctl" + serviceName + "restart")
	if err != nil {
		return err
	}

	return nil
}
