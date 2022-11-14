package usecase

import "CRUD-Chat-Test-Task/internal/core/entity"

func (usecase *usecase) AddMessage(message entity.Message) error {
	err := usecase.chatRepository.AddMessage(&message)
	if err != nil {
		return err
	}
	return nil
}
