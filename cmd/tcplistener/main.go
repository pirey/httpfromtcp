package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer func() {
			fmt.Printf("Connection closed\n")
			f.Close()
			close(ch)
		}()

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
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listening tcp: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accepting connection: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Connection accepted\n")

		ch := getLinesChannel(conn)

		for v := range ch {
			fmt.Printf("%s\n", v)
		}
	}
}
