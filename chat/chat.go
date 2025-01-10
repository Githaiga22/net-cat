// chat.go
package chat

import (
	"fmt"
	"net"
)

func broadcastMessage(message string, clients map[net.Conn]string) {
	for client := range clients {
		fmt.Fprintf(client, message)
	}
}
