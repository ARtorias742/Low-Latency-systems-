package server

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ARtorias742/low_latency_chat/internal/actors"
	"github.com/ARtorias742/low_latency_chat/internal/models"
)

type Server struct {
	room *actors.RoomActor
}

func NewServer(room *actors.RoomActor) *Server {
	return &Server{room: room}
}

func (s *Server) Start() {
	fmt.Println("Type message to send (e.g., 'user1: Hello'):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		var sender, content string
		fmt.Sscanf(input, "%s: %s", &sender, &content)
		s.room.Send(models.Message{Sender: sender, Content: content})
	}
}
