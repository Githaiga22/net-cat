package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"netcat/server"
)


func main() {
	// Default port if not specified
	port := 8989

	// Parse command-line arguments
	if len(os.Args) > 1 {
		p, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("[USAGE]: ./TCPChat $port")
			return
		}
		port = p
	}

	// Start the server
	fmt.Printf("Listening on the port :%d\n", port)
	server.StartServer(port)
}
