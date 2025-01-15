// server.go
package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"netcat/chat"
	"sync"
	"time"
)

var (
	clients    = make(map[net.Conn]string)  // Store client connections with their names
	clientsMux = sync.Mutex{}
)

// messageHistory stores chat history that can be sent to new clients
var messageHistory []string

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Send the welcome message and the ASCII art
	icon, _ := chat.OutputIcon()
	fmt.Fprintf(conn, "\nWelcome to TCP-Chat!\n")
	fmt.Fprintf(conn, "%s\n", string(icon))
	fmt.Fprintf(conn, "[ENTER YOUR NAME]: ")

	// Get the client's name
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	clientName := scanner.Text()

	// Ensure the name is not empty
	if clientName == "" {
		fmt.Fprintf(conn, "Name cannot be empty.\n")
		return
	}

	// Add client to map
	clientsMux.Lock()
	clients[conn] = clientName
	clientsMux.Unlock()

	// Inform other clients that a new client has joined
	broadcastMessage(fmt.Sprintf("%s has joined the chat...\n", clientName))

	// Send previous messages to the client
	for _, msg := range messageHistory {
		fmt.Fprintf(conn, msg)
	}

	// Continuously handle messages from the client
	for {
		if !scanner.Scan() {
			break // Exit loop if the client disconnects
		}

		message := scanner.Text()

		// If message is empty, ignore
		if message == "" {
			continue
		}

		// Format and broadcast message with timestamp
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("[%s][%s]: %s\n", timestamp, clientName, message)
		messageHistory = append(messageHistory, formattedMessage) // Save the message

		// Broadcast the message to all connected clients
		broadcastMessage(formattedMessage)
	}

	// Handle client leaving
	clientsMux.Lock()
	delete(clients, conn)
	clientsMux.Unlock()

	// Broadcast client leaving
	broadcastMessage(fmt.Sprintf("%s has left the chat...\n", clientName))
}

func broadcastMessage(message string) {
	clientsMux.Lock()
	defer clientsMux.Unlock()
	for client := range clients {
		fmt.Fprintf(client, message)
	}
}

func StartServer(port string) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer ln.Close()


	// Accept new client connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn) // Handle client connection in a goroutine
	}
}
