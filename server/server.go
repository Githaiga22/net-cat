// server.go
package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	clients    = make(map[net.Conn]string)  // Store client connections with their names
	clientsMux = sync.Mutex{}
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "\nWelcome to TCP-Chat!\n")
	fmt.Fprintf(conn, "         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n")
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

	// Handle receiving messages
	for scanner.Scan() {
		message := scanner.Text()

		// If message is empty, ignore
		if message == "" {
			continue
		}

		// Format and broadcast message
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("[%s][%s]: %s\n", timestamp, clientName, message)
		messageHistory = append(messageHistory, formattedMessage) // Save the message

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

var messageHistory []string

func startServer(port string) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer ln.Close()

	fmt.Println("Listening on port:", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn) // Handle client connection in a goroutine
	}
}
