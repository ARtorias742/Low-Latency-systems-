package actors

import (
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

func (r *UserActor) Run() {
	for msg := range r.mailbox {
		r.room.Send(models.Message{Sender: r.id, Content: msg.Content})
		r.persistence.Log(models.Message{Sender: r.id, Content: msg.Content})
	}
}

func (r *UserActor) Send(msg models.Message) {
	r.mailbox <- msg
}
