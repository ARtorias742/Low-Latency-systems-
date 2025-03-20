package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Get username (from args or prompt)
	var username string
	if len(os.Args) == 2 {
		username = os.Args[1]
	} else {
		fmt.Print("Enter your username: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			username = strings.TrimSpace(scanner.Text())
		}
		if username == "" {
			fmt.Println("Username cannot be empty. Exiting.")
			os.Exit(1)
		}
	}

	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Send username to server
	fmt.Fprintf(conn, "%s\n", username)
	fmt.Printf("Connected as %s. Type messages to chat (or 'exit' to quit):\n", username)

	// Goroutine to read server messages
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Connection lost: %v\n", err)
		}
	}()

	// Send messages from user input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := strings.TrimSpace(scanner.Text())
		if msg == "exit" {
			fmt.Println("Disconnecting...")
			break
		}
		if msg != "" {
			fmt.Fprintf(conn, "%s\n", msg)
		}
	}
}
