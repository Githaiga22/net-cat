package client

// Handles the client logic (connecting, sending, receiving messages)

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartClient(serverAddress string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverAddress, port))
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	// Handle greeting and name input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to TCP-Chat!")
	fmt.Print("[ENTER YOUR NAME]: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Name cannot be empty. Exiting.")
		return
	}

	// Send the name to the server
	fmt.Fprintf(conn, "%s\n", name)

	// Listen for incoming messages from the server
	go receiveMessages(conn)

	// Sending messages to the server
	for {
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "" {
			continue
		}

		// Send the message to the server
		fmt.Fprintf(conn, "%s\n", message)
	}
}

// Receive messages from the server
func receiveMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, _ := reader.ReadString('\n')
		fmt.Print(message)
	}
}
