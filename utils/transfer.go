package utils

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func SendFile(address string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	fileName := fileInfo.Name()
	fileSize := fileInfo.Size()

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Printf("Sending file: %s...\n", fileName)

	err = binary.Write(conn, binary.BigEndian, uint64(len(fileName)))
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(fileName))
	if err != nil {
		return err
	}

	err = binary.Write(conn, binary.BigEndian, uint64(fileSize))
	if err != nil {
		return err
	}

	_, err = io.Copy(conn, file)
	if err != nil {
		return err
	}

	buf := make([]byte, 2)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return fmt.Errorf("failed to get confirmation: %v", err)
	}

	confirmation := string(buf)
	if confirmation != "OK" {
		return fmt.Errorf("receiver sent unexpected status: %s", confirmation)
	}

	fmt.Println("File sent and confirmed by receiver!")
	return nil
}

func ReceiveFile(listenAddr string) error {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Println("Waiting for connection...")

	conn, err := listener.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	var nameLen uint64
	err = binary.Read(conn, binary.BigEndian, &nameLen)
	if err != nil {
		return err
	}

	nameBuf := make([]byte, nameLen)
	_, err = io.ReadFull(conn, nameBuf)
	if err != nil {
		return err
	}
	fileName := string(nameBuf)

	fileName = filepath.Base(fileName)

	var fileSize uint64
	err = binary.Read(conn, binary.BigEndian, &fileSize)
	if err != nil {
		return err
	}

	fmt.Println("Receiving file...")

	outputFile, err := os.Create("received_" + fileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.CopyN(outputFile, conn, int64(fileSize))
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte("OK"))
	if err != nil {
		return err
	}

	fmt.Printf("File saved as received_%s!\n", fileName)
	return nil
}
