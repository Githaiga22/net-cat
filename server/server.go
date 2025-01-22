// server.go
package server

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "sync"
    "time"

    "netcat/chat"
)

type Message struct {
    sender  net.Conn
    content string
}

var (
    clients       = make(map[net.Conn]string) // Store client connections with their names
    clientsMux    = sync.Mutex{}
    maxClients    = 10
    activeClients int
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
    clientName := ""

    for {
        if !scanner.Scan() {
            return
        }
        clientName = scanner.Text()

        // Ensure the name is not empty
        if clientName == "" {
            fmt.Fprintf(conn, "Name cannot be empty. Please enter a valid name.\n")
            fmt.Fprintf(conn, "[ENTER YOUR NAME]: ")
            continue
        }

        if !IsNameUnique(clientName) {
            fmt.Fprintf(conn, "Name already taken. Choose another name.\n")
            fmt.Fprintf(conn, "[ENTER YOUR NAME]: ")
            continue
        } else {
            // Add client to map
            clientsMux.Lock()
            clients[conn] = clientName
            clientsMux.Unlock()
            break
        }
    }

    // Send previous messages to the client
    for _, msg := range messageHistory {
        fmt.Fprintf(conn, "%s", msg)
    }

    // Broadcast message that the client has joined
    broadcastMessage(fmt.Sprintf("%s has joined the chat\n", clientName), conn)

    // Continuously handle messages from the client
    for {
		clientName = clients[conn]		
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
		conn.Write([]byte("\033[F\033[K"))
        messageHistory = append(messageHistory, formattedMessage) // Save the message
        conn.Write([]byte(formattedMessage))
        // Broadcast the message to all connected clients
        broadcastMessage(formattedMessage, conn)
    }

    // Handle client leaving with a timestamp
    clientsMux.Lock()
    delete(clients, conn)
    activeClients--
    clientsMux.Unlock()

    // Broadcast client leaving with timestamp
    broadcastMessage(fmt.Sprintf("%s has left the chat\n", clientName), conn)
}

func broadcastMessage(message string, excludeConn net.Conn) {
    clientsMux.Lock()
    defer clientsMux.Unlock()

    for client := range clients {
        if excludeConn != client {
            fmt.Fprintf(client, "%s", message)
        }
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

        clientsMux.Lock()
        if activeClients >= maxClients {
            clientsMux.Unlock()
            conn.Close()
            log.Println("Connection refused: Max clients reached")
            continue
        }
        activeClients++
        clientsMux.Unlock()

        go handleClient(conn) // Handle client connection in a goroutine
    }
}

func IsNameUnique(clientName string) bool {
    clientsMux.Lock()
    defer clientsMux.Unlock()

    for _, name := range clients {
        if name == clientName {
            return false
        }
    }
    return true
}
