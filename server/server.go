package server


// Handles the server side logic ~ broadcasting and multiple clients

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
	"netcat/chat"
)

// Global variable to store all connected clients
var clients = make(map[net.Conn]string)
var mu sync.Mutex

// Start the server and listen for incoming connections
func StartServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Server main loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Spawn a new goroutine to handle each client
		go handleClient(conn)
	}
}

// Handle individual client connections
func handleClient(conn net.Conn) {
	// Greet the client
	fmt.Fprintln(conn, "Welcome to TCP-Chat!")
	fmt.Fprintln(conn, "         _nnnn_")
	fmt.Fprintln(conn, "        dGGGGMMb")
	// ... Display ASCII art

	// Ask for the name of the client
	fmt.Fprint(conn, "[ENTER YOUR NAME]: ")
	reader := bufio.NewReader(conn)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Validate name
	if name == "" {
		fmt.Fprintln(conn, "Name cannot be empty. Disconnecting.")
		conn.Close()
		return
	}

	// Add the client to the chat
	mu.Lock()
	clients[conn] = name
	mu.Unlock()

	// Notify other clients that a new client has joined
	chat.Broadcast(fmt.Sprintf("%s has joined our chat...", name), conn)

	// Send previous chat messages to the new client
	chat.SendHistory(conn)

	// Listen for messages from this client
	for {
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		// Exit condition if client sends an empty message or disconnects
		if message == "" {
			break
		}

		// Broadcast message to all other clients
		chat.Broadcast(fmt.Sprintf("[%s]: %s", name, message), conn)
	}

	// Handle client exit
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
	chat.Broadcast(fmt.Sprintf("%s has left our chat...", name), nil)

	conn.Close()
}
