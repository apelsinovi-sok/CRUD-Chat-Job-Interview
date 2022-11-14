package model

import (
	"github.com/google/uuid"
)

type Message struct {
	ID          uuid.UUID
	AuthorID    uuid.UUID
	ChatID      uuid.UUID
	CreatedAt   string
	MessageText string
}
