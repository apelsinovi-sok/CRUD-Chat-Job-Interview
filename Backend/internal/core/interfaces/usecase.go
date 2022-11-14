package interfaces

import "CRUD-Chat-Test-Task/internal/core/entity"

type Usecase interface {
	CreateChat(chat entity.Chat) error
	AddMessage(message entity.Message) error
	GetListMessages(chatName string, limit int) (entity.ListIdMessages, error)
	GetMessage(id string) (entity.Message, error)
}
