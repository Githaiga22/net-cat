package utils

import (
	"fmt"
	"time"
)

// Format a message with a timestamp and user name
func FormatMessage(name, message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s][%s]: %s", timestamp, name, message)
}
