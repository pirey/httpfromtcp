package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		fmt.Fprintf(os.Stderr, "resolve udp: %v\n", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "diap udp: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	rd := os.Stdin
	buf := bufio.NewReader(rd)

	for {
		fmt.Printf(">")
		s, err := buf.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "buf read stdin: %v\n", err)
			os.Exit(1)
		}

		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Fprintf(os.Stderr, "buf write: %v\n", err)
			continue
		}
	}
}
