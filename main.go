// main.go
package main

import (
	"fmt"
	"os"
	// "netcat/client"
	"netcat/server"
	"netcat/utils"
)

func main() {
	args := os.Args

	// Case 1: If no arguments are provided, start the server on port 8989
	if len(args) == 1 {
		// Default port 8989 if no port specified
		fmt.Println("Listening on the port :8989")
		server.StartServer("8989")
		return
	}

	// Case 2: If exactly 1 argument (port) is provided, start the server with that port
	if len(args) == 2 {
		// Server mode: Start server on the specified port
		fmt.Println("Listening on the port:", args[1])
		server.StartServer(args[1])
		return
	}

	// Case 3: If two arguments are provided (port and address), it's invalid usage
	if len(args) == 3 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	
	// Utility function to clear the terminal screen (it needs to be called, not imported globally)
	utils.ClearTerminal()
}
