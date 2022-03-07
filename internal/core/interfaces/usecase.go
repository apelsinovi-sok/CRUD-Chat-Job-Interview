package interfaces

import "CRUD-Chat-Test-Task/internal/core/entity"

type Usecase interface {
	CreateChat(chat entity.Chat) error
}
