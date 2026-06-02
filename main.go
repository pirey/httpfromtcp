package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

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
				ch <- p
			}

			line = parts[len(parts)-1]
		}
	}()

	return ch
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	ch := getLinesChannel(f)

	for v := range ch {
		fmt.Printf("read: %s\n", v)
	}
}
