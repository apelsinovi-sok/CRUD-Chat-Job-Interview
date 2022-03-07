package main

import (
	"CRUD-Chat-Test-Task/infrastructure/DB"
	"CRUD-Chat-Test-Task/internal/core/usecase"
	"CRUD-Chat-Test-Task/internal/delivery/http"
	"CRUD-Chat-Test-Task/internal/repository"
)

func main() {
	db := DB.New()
	chatRepository := repository.NewChatRepository(db)
	useCase := usecase.New(chatRepository)
	httpServer := http.New(useCase)
	httpServer.Run()
}
