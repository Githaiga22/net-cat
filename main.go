// main.go
package main

import (
	"fmt"
	"os"
	"netcat/client"
	"netcat/server"
	"netcat/utils"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		// Default port 8989 if no port specified
		server.StartServer("8989")
		return
	}

	if len(args) == 2 {
		// Server mode
		server.StartServer(args[1])
	} else if len(args) == 3 {
		// Client mode
		serverAddr := args[1]
		port := args[2]
		client.StartClient(serverAddr, port)
	} else {
		// Invalid usage
		fmt.Println("[USAGE]: ./TCPChat $port")
	}
	
	// Utility function to clear the terminal screen (it needs to be called, not imported globally)
	utils.ClearTerminal()
}
