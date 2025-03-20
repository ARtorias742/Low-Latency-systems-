package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run cmd/client/main.go <username>")
		os.Exit(1)
	}
	username := os.Args[1]

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Send username to server
	fmt.Fprintf(conn, "%s\n", username)

	// Goroutine to read server messages
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// Read user input and send to server
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "exit" {
			break
		}
		fmt.Fprintf(conn, "%s\n", msg)
	}
}
