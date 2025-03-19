package main

import (
	"github.com/ARtorias742/low_latency_chat/internal/actors"
	"github.com/ARtorias742/low_latency_chat/internal/server"
)

func main() {
	// Create a room actor to manage broadcasting
	room := actors.NewRoomActor()

	// Start the persistence actor for async logging
	persistence := actors.NewPersistenceActor()
	go persistence.Run()

	// Simulate two users
	user1 := actors.NewUserActor("user1", room, persistence)
	user2 := actors.NewUserActor("user2", room, persistence)
	go user1.Run()
	go user2.Run()

	// Start a simple server to accept messages
	srv := server.NewServer(room)
	srv.Start()
}
