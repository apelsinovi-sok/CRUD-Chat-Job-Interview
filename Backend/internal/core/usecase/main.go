package usecase

import "CRUD-Chat-Test-Task/internal/core/interfaces"

type usecase struct {
	chatRepository interfaces.ChatRepository
}

func New(chatRepository interfaces.ChatRepository) interfaces.Usecase {
	return &usecase{chatRepository: chatRepository}
}
