package usecase

import "CRUD-Chat-Test-Task/internal/core/interfaces"

type useCase struct {
	chatRepository interfaces.ChatRepository
}

func New(chatRepository interfaces.ChatRepository) interfaces.Usecase {
	return &useCase{chatRepository: chatRepository}
}
