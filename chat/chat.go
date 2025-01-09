package chat


// Contains shared chat logic (e.g., storing messages, managing users)

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var messageHistory []string
var mu sync.Mutex

// Broadcast a message to all connected clients
func Broadcast(message string, sender net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	// Add timestamp to the message
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formattedMessage := fmt.Sprintf("[%s]: %s\n", timestamp, message)

	// Add the message to the history
	messageHistory = append(messageHistory, formattedMessage)

	// Send the message to all clients
	for conn := range clients {
		if conn != sender {
			_, err := conn.Write([]byte(formattedMessage))
			if err != nil {
				// Handle error
				fmt.Println("Error sending message to client:", err)
			}
		}
	}
}

// Send the message history to the new client
func SendHistory(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	for _, msg := range messageHistory {
		conn.Write([]byte(msg))
	}
}
