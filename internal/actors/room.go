package actors

import "github.com/ARtorias742/low_latency_chat/internal/models"

type RoomActor struct {
	users   map[string]*UserActor
	mailbox chan models.Message
}

func NewRoomActor() *RoomActor {
	return &RoomActor{
		users:   make(map[string]*UserActor),
		mailbox: make(chan models.Message, 100),
	}
}

func (r *RoomActor) Run() {
	for msg := range r.mailbox {
		// Broadcast to all users except the sender
		for id, user := range r.users {
			if id != msg.Sender {
				user.Send(msg)
			}
		}
	}
}

func (r *RoomActor) Send(msg models.Message) {
	r.mailbox <- msg
}

func (r *RoomActor) AddUser(user *UserActor) {
	r.users[user.id] = user
}
