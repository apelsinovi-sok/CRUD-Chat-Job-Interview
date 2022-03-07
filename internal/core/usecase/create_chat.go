package usecase

import "CRUD-Chat-Test-Task/internal/core/entity"

func (usecase *usecase) CreateChat(chat entity.Chat) error {
	err := usecase.chatRepository.CreateChat(&chat)
	if err != nil {
		return err
	}
	return nil
}
