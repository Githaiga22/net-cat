// utils.go
package chat

import (
	"fmt"
	"os"
	"os/exec"
)

func handleError(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
		os.Exit(1)
	}
}


// Utility function to clear the terminal screen
func clearTerminal() {
	cmd := exec.Command("clear") // or "cls" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}