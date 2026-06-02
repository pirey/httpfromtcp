package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("Failed to open file\n")
		return
	}
	line := ""
	b := make([]byte, 8)
	for {
		n, err := f.Read(b)
		if err == io.EOF {
			break
		}

		line += string(b[:n])
		parts := strings.Split(line, "\n")
		for _, p := range parts[:len(parts)-1] {
			fmt.Printf("read: %s\n", p)
		}

		line = parts[len(parts)-1]
	}
}
