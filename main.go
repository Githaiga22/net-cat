package main

import (
	"fmt"
	"os"
	"os/signal"
	"netcat/server" // Import the server package
)

func main() {
	// Default port (or use a port provided as an argument)
	port := "8989"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Graceful shutdown handling (e.g., Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt) // Listen for interrupt signal (Ctrl+C)

	// Start the server
	fmt.Printf("Listening on port: %s\n", port)
	go server.StartServer(port) // Run the server in a separate goroutine

	// Wait for shutdown signal (Ctrl+C)
	<-sigChan
	fmt.Println("\nServer shutting down...")
}
