package interfaces

import "CRUD-Chat-Test-Task/internal/core/entity"

type ChatRepository interface {
	CreateChat(chat *entity.Chat) error
	AddMessage(message *entity.Message) error
}
