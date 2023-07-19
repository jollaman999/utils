package systemctl

import (
	"github.com/jollaman999/utils/cmd"
	"github.com/jollaman999/utils/syscheck"
)

func controlServiceCommon(serviceName string, option string) error {
	err := syscheck.CheckRoot()
	if err != nil {
		return err
	}

	err = cmd.RunCMD("systemctl" + serviceName + option)
	if err != nil {
		return err
	}

	return nil
}

// StartService : Start the service with systemctl command.
func StartService(serviceName string) error {
	return controlServiceCommon(serviceName, "start")
}

// StopService : Stop the service with systemctl command.
func StopService(serviceName string) error {
	return controlServiceCommon(serviceName, "stop")
}

// RestartService : Restart the service with systemctl command.
func RestartService(serviceName string) error {
	return controlServiceCommon(serviceName, "restart")
}

// CustomOption : Run systemctl command with the custom option.
func CustomOption(serviceName string, option string) error {
	return controlServiceCommon(serviceName, option)
}
