package chat

import (
	"fmt"
	"os"
)

func OutputIcon() ([]byte, error) {
	content, err := os.ReadFile("docs/icon.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return nil, err
	}

	// fmt.Println(string(content))

	return content, nil
}
