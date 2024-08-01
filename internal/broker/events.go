package broker

import "testkafka/internal/models"

type MessageEvent struct {
	ID      int
	Message string
}

func NewMessageEvent(message models.Message) MessageEvent {
	return MessageEvent{
		ID:      message.ID,
		Message: message.Message,
	}
}
