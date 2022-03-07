package usecase

import "CRUD-Chat-Test-Task/internal/core/entity"

func (usecase *usecase) GetMessage(id int) (entity.Message, error) {
	message, err := usecase.chatRepository.GetMessage(id)
	if err != nil {
		return entity.Message{}, err
	}
	return message, nil
}
