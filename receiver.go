package main

import (
	"fmt"
	"net"
)

func startReceiver() {
	addr, err := net.ResolveUDPAddr("udp", ":10609")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	n, remoteAddress, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	text := string(buf[:n])
	fmt.Println(remoteAddress.String(), "|", text)
}
