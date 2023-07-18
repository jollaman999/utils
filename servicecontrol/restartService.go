package servicecontrol

import (
	"github.com/jollaman999/utils/cmd"
)

// RestartService : Restart the service with systemctl command.
func RestartService(serviceName string) error {
	err := cmd.RunCMD("systemctl" + serviceName + "restart")
	if err != nil {
		return err
	}

	return nil
}
