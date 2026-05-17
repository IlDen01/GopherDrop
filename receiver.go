package main

import (
	"GopherDrop/protocol"
	"GopherDrop/utils"
	"encoding/json"
	"fmt"
	"net"
)

func startReceiver() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":10609")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Waiting for sender...")

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var msg protocol.Message
	err = json.Unmarshal(buf[:n], &msg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if msg.Type != "announce" {
		fmt.Printf("Ignored message type: %s\n", msg.Type)
		return
	}

	fmt.Printf("Sender found: %s\n", msg.Name)

	myIP, err := utils.GetOutboundIP()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resp := protocol.Message{
		Type: "response",
		Name: deviceName,
		IP:   myIP,
		Port: 10610,
	}

	respData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = conn.WriteToUDP(respData, remoteAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	errChan := make(chan error)

	go func() {
		result := utils.ReceiveFile(":10610")
		errChan <- result
	}()

	if err := <-errChan; err != nil {
		fmt.Println("Transfer error:", err)
	}
}
