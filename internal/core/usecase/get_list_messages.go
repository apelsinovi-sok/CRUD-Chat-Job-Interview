package usecase

import "CRUD-Chat-Test-Task/internal/core/entity"

func (usecase *usecase) GetListMessages(chatName string, limit int) (entity.ListIdMessages, error) {
	ListIdMessages, err := usecase.chatRepository.GetListMessages(chatName, limit)
	if err != nil {
		return nil, err
	}
	return ListIdMessages, nil
}
