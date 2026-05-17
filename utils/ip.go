package utils

import (
	"net"
)

func GetOutboundIP() (string, error) {
	connection, connectionError := net.Dial("udp", "8.8.8.8:80")
	if connectionError != nil {
		return "", connectionError
	}
	defer connection.Close()

	localAddress := connection.LocalAddr().(*net.UDPAddr)

	return localAddress.IP.String(), nil
}
