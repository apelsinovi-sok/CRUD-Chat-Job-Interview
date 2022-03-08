package entity

import "github.com/google/uuid"

type Message struct {
	ID            int
	ChatID        uuid.UUID
	MessageAuthor string `json:"message_author" binding:"required"`
	ChatName      string `json:"chat_name" binding:"required"`
	MessageText   string `json:"message_text" binding:"required"`
}
