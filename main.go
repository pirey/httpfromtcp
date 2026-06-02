package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("Failed to open file\n")
		return
	}
	b := make([]byte, 8)
	for {
		n, err := f.Read(b)
		if n == 0 || err == io.EOF {
			break
		}
		fmt.Printf("read: %s\n", b)
	}
}
