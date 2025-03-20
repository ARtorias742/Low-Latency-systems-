package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/ARtorias742/low_latency_chat/internal/actors"
	"github.com/ARtorias742/low_latency_chat/internal/models"
)

type Server struct {
	room        *actors.RoomActor
	persistence *actors.PersistenceActor
	addr        string
}

func NewServer(room *actors.RoomActor, persistence *actors.PersistenceActor, addr string) *Server {
	return &Server{
		room:        room,
		persistence: persistence,
		addr:        addr,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server listening for connections...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read username from client
	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		fmt.Println("Client disconnected before sending username")
		return
	}
	username := scanner.Text()
	if username == "" {
		fmt.Println("Received empty username; closing connection")
		return
	}

	// Create and register UserActor with connection
	user := actors.NewUserActor(username, s.room, s.persistence, conn)
	s.room.AddUser(user)
	go user.Run()

	fmt.Printf("%s connected\n", username)

	// Handle incoming messages
	for scanner.Scan() {
		content := scanner.Text()
		if content != "" {
			user.Send(models.Message{Sender: username, Content: content})
		}
	}
	fmt.Printf("%s disconnected\n", username)
}
