package entity

import "github.com/google/uuid"

type Chat struct {
	ID         uuid.UUID
	ChatName   string `json:"chat_name" binding:"required"`
	ChatAuthor string `json:"chat_author" binding:"required"`
	Messages   []Message
}
