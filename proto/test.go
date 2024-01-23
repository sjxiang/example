package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println(LocalIP())

	resp, _ := GetOutBoundIP()
	fmt.Println(resp)
}

func LocalIP() string {
	ipList := []string{"114.114.114.114:80", "8.8.8.8:80"}
	for _, ip := range ipList {
		conn, err := net.Dial("udp", ip)
		if err != nil {
			continue
		}
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		conn.Close()
		return localAddr.IP.String()
	}

	return ""
}


func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
