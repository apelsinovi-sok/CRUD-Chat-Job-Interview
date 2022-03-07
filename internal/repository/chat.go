package repository

import (
	"CRUD-Chat-Test-Task/internal/core/entity"
	"CRUD-Chat-Test-Task/internal/core/interfaces"
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
