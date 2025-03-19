package actors

import (
	"fmt"

	"github.com/ARtorias742/low_latency_chat/internal/models"
)

type PersistenceActor struct {
	mailbox chan models.Message
}

func NewPersistenceActor() *PersistenceActor {
	return &PersistenceActor{
		mailbox: make(chan models.Message, 100),
	}
}

func (p *PersistenceActor) Run() {
	for msg := range p.mailbox {
		// Simulate async DB write
		fmt.Printf("Logged: %s - %s\n", msg.Sender, msg.Content)
	}
}

func (p *PersistenceActor) Log(msg models.Message) {
	p.mailbox <- msg
}
