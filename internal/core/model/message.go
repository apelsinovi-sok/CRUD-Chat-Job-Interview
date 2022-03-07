package model

import "github.com/google/uuid"

type Message struct {
	ID          int
	AuthorID    uuid.UUID
	ChatID      uuid.UUID
	MessageText string
}
