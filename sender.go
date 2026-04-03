package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func startSender() {
	fmt.Println("Starting sender...")
	path, err := getPath()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	content, err := readFirstChunk(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Size:", len(content), "| Content:", content)
	content, err = readFullChunk(path)
	if err != nil {
		fmt.Print("Error: ", err, " | ")
	}
	fmt.Println("Size:", len(content), "| Content:", content)
}

func getPath() (string, error) {
	fmt.Println("Enter path:")
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

func readFirstChunk(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buffer := make([]byte, 16)
	size, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buffer[:size], nil
}

func readFullChunk(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buffer := make([]byte, 16)
	size, err := io.ReadFull(file, buffer)
	if err != nil {
		return buffer[:size], err
	}
	return buffer[:size], nil
}
