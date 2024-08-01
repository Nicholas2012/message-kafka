package models

import "time"

type Message struct {
	ID        int
	Message   string
	CreatedAt time.Time
	SentAt    *time.Time
}

func (m *Message) SetSentNow() {
	now := time.Now()
	m.SentAt = &now
}
