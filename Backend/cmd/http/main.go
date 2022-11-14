package main

import (
	"CRUD-Chat-Test-Task/infrastructure/DB"
	"CRUD-Chat-Test-Task/infrastructure/configuration"
	"CRUD-Chat-Test-Task/internal/core/usecase"
	"CRUD-Chat-Test-Task/internal/delivery/http"
	"CRUD-Chat-Test-Task/internal/repository"
)

func main() {
	config := configuration.New(".")
	db := DB.New(config.DatabaseUrl)
	chatRepository := repository.NewChatRepository(db)
	useCase := usecase.New(chatRepository)
	httpServer := http.New(useCase)
	httpServer.Run(config.Port)
}
