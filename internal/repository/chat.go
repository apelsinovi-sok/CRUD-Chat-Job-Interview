package repository

import (
	"CRUD-Chat-Test-Task/internal/core/entity"
	"CRUD-Chat-Test-Task/internal/core/interfaces"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) interfaces.ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) CreateChat(chat *entity.Chat) error {
	db := r.db.Create(chat)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (r chatRepository) AddMessage(message *entity.Message) error {
	chat := entity.Chat{}

	db := r.db.Where(&entity.Chat{ChatName: message.ChatName}).First(&chat)
	if db.Error != nil {
		return errors.New("chat not found")
	}
	message.ChatID = chat.ID

	fmt.Println(chat)
	fmt.Println(*message)

	db = r.db.Create(message)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
