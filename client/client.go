// client.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func startClient(serverAddr, port string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", serverAddr, port))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Receive and print welcome message
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "[ENTER YOUR NAME]:" {
			break
		}
	}

	// Read user name input
	fmt.Print("[ENTER YOUR NAME]: ")
	var name string
	fmt.Scanln(&name)

	// Send name to server
	fmt.Fprintf(conn, "%s\n", name)

	// Handle sending and receiving messages
	go receiveMessages(conn)

	// Read and send messages from user
	for {
		var message string
		fmt.Scanln(&message)
		if message != "" {
			fmt.Fprintf(conn, "%s\n", message)
		}
	}
}

func receiveMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
