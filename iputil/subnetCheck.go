package iputil

import (
	"net"
)

func IsPrivateAClass(IP net.IP) bool {
	if IP[0] == 10 {
		return true
	}

	return false
}

func IsPrivateBClass(IP net.IP) bool {
	if IP[0] == 172 &&
		(IP[1] >= 16 && IP[1] <= 31) {
		return true
	}

	return false
}

func IsPrivateCClass(IP net.IP) bool {
	if IP[0] == 192 && IP[1] == 168 {
		return true
	}

	return false
}
