package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var deviceName string

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		deviceName = "UnknownDevice"
	} else {
		deviceName = hostname
	}
}

func printBanner() {
	fmt.Print(`
   ____             _               ____                  
  / ___| ___  _ __ | |__   ___ _ __|  _ \ _ __ ___  _ __  
 | |  _ / _ \| '_ \| '_ \ / _ \ '__| | | | '__/ _ \| '_ \ 
 | |_| | (_) | |_) | | | |  __/ |  | |_| | | | (_) | |_) |
  \____|\___/| .__/|_| |_|\___|_|  |____/|_|  \___/| .__/ 
             |_|                                    |_|    
`)
	fmt.Printf(" Device Name: %s\n", deviceName)
}

func printMenu() {
	fmt.Println("\nSelect mode:")
	fmt.Println("1) Send file")
	fmt.Println("2) Receive file")
	fmt.Println("3) Change device name")
	fmt.Println("0) Exit")
}

func changeName() {
	fmt.Print("Enter new name: ")
	reader := bufio.NewReader(os.Stdin)
	newName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	newName = strings.TrimSpace(newName)
	if newName != "" {
		deviceName = newName
		fmt.Println("Name changed to:", deviceName)
	}
}

func modeChoice() {
	for {
		printMenu()
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch choice {
		case 1:
			startSender()
		case 2:
			startReceiver()
		case 3:
			changeName()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func main() {
	printBanner()
	modeChoice()
}
