// client.go
package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartClient(serverAddr, port string) {
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

	// Start receiving messages concurrently
	go receiveMessages(conn)

	// Handle sending and receiving messages
	for {
		var message string
		fmt.Scanln(&message)
		if message != "" {
			fmt.Fprintf(conn, "%s\n", message) // Send message to server
		}
	}
}
func receiveMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Print received messages
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading from server:", scanner.Err())
	}
}
