package actors

import (
	"fmt"
	"net"

	"github.com/ARtorias742/low_latency_chat/internal/models"
)

type UserActor struct {
	id          string
	room        *RoomActor
	persistence *PersistenceActor
	mailbox     chan models.Message
	conn        net.Conn // Add connection to send messages to client
}

func NewUserActor(id string, room *RoomActor, persistence *PersistenceActor, conn net.Conn) *UserActor {
	return &UserActor{
		id:          id,
		room:        room,
		persistence: persistence,
		mailbox:     make(chan models.Message, 100),
		conn:        conn,
	}
}

func (u *UserActor) Run() {
	for msg := range u.mailbox {
		// Send received message to client, don’t forward to room
		if u.conn != nil && msg.Sender != u.id { // Only send if not from self
			fmt.Fprintf(u.conn, "%s: %s\n", msg.Sender, msg.Content)
		}
	}
}

func (u *UserActor) Send(msg models.Message) {
	u.mailbox <- msg
	// Only forward original messages from this user to room and persistence
	if msg.Sender == u.id {
		u.room.Send(msg)
		u.persistence.Log(msg)
	}
}
