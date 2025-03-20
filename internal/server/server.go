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

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
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
		return
	}

	username := scanner.Text()

	// create and start UserActor
	user := actors.NewUserActor(username, s.room, s.persistence)
	s.room.AddUser(user)
	go user.Run()

	// Send message from client to UserActor
	for scanner.Scan() {
		content := scanner.Text()
		user.Send(models.Message{Sender: username, Content: content})
	}
}
