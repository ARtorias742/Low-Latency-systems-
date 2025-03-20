package main

import (
	"fmt"

	"github.com/ARtorias742/low_latency_chat/internal/actors"
	"github.com/ARtorias742/low_latency_chat/internal/server"
)

func main() {
	room := actors.NewRoomActor()
	go room.Run()

	persistence := actors.NewPersistenceActor()
	go persistence.Run()

	srv := server.NewServer(room, persistence, ":8080")
	fmt.Println("Server running on :8080")
	srv.Start()
}
