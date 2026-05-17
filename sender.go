package main

import (
	"GopherDrop/protocol"
	"GopherDrop/utils"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func startSender() {
	path, err := getPath()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	myIP, err := utils.GetOutboundIP()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	addr := &net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 10609,
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 0})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	msg := protocol.Message{
		Type: "announce",
		Name: deviceName,
		IP:   myIP,
		Port: 0,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = conn.WriteToUDP(data, addr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Looking for receivers...")

	var receivers []protocol.Message
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			var netErr net.Error
			if errors.As(err, &netErr) && netErr.Timeout() {
				break
			}
			fmt.Println("Read error:", err)
			break
		}

		var resp protocol.Message
		err = json.Unmarshal(buf[:n], &resp)
		if err == nil {
			if resp.Type == "response" {
				receivers = append(receivers, resp)
			}
		}
	}

	if len(receivers) == 0 {
		fmt.Println("No receivers found.")
		return
	}

	fmt.Println("\nFound devices:")
	for i, r := range receivers {
		fmt.Printf("%d) %s (%s:%d)\n", i+1, r.Name, r.IP, r.Port)
	}

	fmt.Print("Select device number: ")
	var choice int
	_, err = fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if choice < 1 || choice > len(receivers) {
		fmt.Println("Invalid selection")
		return
	}

	selected := receivers[choice-1]
	targetAddr := selected.IP + ":" + strconv.Itoa(selected.Port)

	err = utils.SendFile(targetAddr, path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func getPath() (string, error) {
	fmt.Print("Enter path to file: ")
	reader := bufio.NewReader(os.Stdin)
	path, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	path = strings.TrimSpace(path)
	if path == "" {
		return "", errors.New("path is empty")
	}
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		return "", errors.New("path is a directory")
	}
	return path, nil
}
