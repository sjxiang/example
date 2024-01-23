package main

import (
	"net"
)


// GetOutboundIP 获取本机的出口 IP 地址
func GetOutboundIP() (string , error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

