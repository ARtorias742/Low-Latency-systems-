package actors

import (
	"fmt"

	"github.com/ARtorias742/low_latency_chat/internal/models"
)

type UserActor struct {
	id          string
	room        *RoomActor
	persistence *PersistenceActor
	mailbox     chan models.Message
}

func NewUserActor(id string, room *RoomActor, persistence *PersistenceActor) *UserActor {
	return &UserActor{
		id:          id,
		room:        room,
		persistence: persistence,
		mailbox:     make(chan models.Message, 100),
	}
}

func (u *UserActor) Run() {
	for msg := range u.mailbox {
		// Process message: forward to room and log asynchronously

		fmt.Printf("%s received: %s\n", u.id, msg.Content)

		u.room.Send(models.Message{Sender: u.id, Content: msg.Content})

		u.persistence.Log(models.Message{Sender: u.id, Content: msg.Content})
	}
}

func (u *UserActor) Send(msg models.Message) {
	u.mailbox <- msg
}
