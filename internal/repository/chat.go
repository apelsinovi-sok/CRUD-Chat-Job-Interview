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

func (r chatRepository) GetListMessages(chatName string, limit int) (entity.ListIdMessages, error) {
	chat := entity.Chat{}
	ListIdMessages := entity.ListIdMessages{}
	db := r.db.Where("chat_name  = ?", chatName).First(&chat)
	if db.Error != nil {
		return nil, errors.New("chat not found")
	}
	db = r.db.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Where("chat_id = ?", chat.ID).Limit(limit).Order("messages.id DESC")
	}).Find(&chat)
	if db.Error != nil {
		return nil, db.Error
	}

	for i, v := range chat.Messages {
		ListIdMessages[i] = v.ID
	}
	fmt.Println(ListIdMessages)
	return ListIdMessages, nil
}
