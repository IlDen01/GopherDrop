package main

import (
	"fmt"
)

func printBanner() {
	fmt.Print(`
   ____             _               ____                  
  / ___| ___  _ __ | |__   ___ _ __|  _ \ _ __ ___  _ __  
 | |  _ / _ \| '_ \| '_ \ / _ \ '__| | | | '__/ _ \| '_ \ 
 | |_| | (_) | |_) | | | |  __/ |  | |_| | | | (_) | |_) |
  \____|\___/| .__/|_| |_|\___|_|  |____/|_|  \___/| .__/ 
             |_|                                    |_|    
`)
}

func printMenu() {
	fmt.Println("Select mode:")
	fmt.Println("1) Send")
	fmt.Println("2) Receive")
}

func modeChoice() {
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Reading error:", err)
		return
	}

	switch choice {
	case 1:
		startSender()
	case 2:
		startReceiver()
	default:
		fmt.Println("Invalid choice")
	}
}

func main() {
	printBanner()
	printMenu()
	modeChoice()
}
