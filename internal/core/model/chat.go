package model

import "github.com/google/uuid"

type Chat struct {
	ID       uuid.UUID
	ChatName string
	AuthorID uuid.UUID
}
