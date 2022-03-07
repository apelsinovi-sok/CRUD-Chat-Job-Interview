package repository

import (
	"CRUD-Chat-Test-Task/internal/core/entity"
	"CRUD-Chat-Test-Task/internal/core/interfaces"
	"CRUD-Chat-Test-Task/internal/core/model"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) interfaces.ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) CreateChat(chatEntity *entity.Chat) error {
	user := model.User{}
	chat := model.Chat{}
	userId := uuid.New()
	//Поиск юзера по нику. Если не нашел, создать юзера
	db := r.db.Where("name = ?", chatEntity.ChatAuthor).Limit(1).First(&user)
	if db.RowsAffected == 0 {
		user.ID = userId
		user.Name = chatEntity.ChatAuthor
		r.db.Create(&user)
	}
	//Записать все данные из сущности в модель, для разбиения сущности на таблицы
	chat.ID = uuid.New()
	chat.AuthorID = userId
	chat.ChatName = chatEntity.ChatName
	db = r.db.Create(&chat)

	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (r *chatRepository) AddMessage(message *entity.Message) error {
	chat := entity.Chat{}

	db := r.db.Where(&entity.Chat{ChatName: message.ChatName}).First(&chat)
	if db.Error != nil {
		return errors.New("chat not found")
	}
	message.ChatID = chat.ID

	db = r.db.Create(message)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (r *chatRepository) GetListMessages(chatName string, limit int) (entity.ListIdMessages, error) {
	chat := entity.Chat{}
	ListIdMessages := entity.ListIdMessages{}

	db := r.db.Where("chat_name  = ?", chatName).First(&chat)

	if db.Error != nil {
		return entity.ListIdMessages{}, errors.New("chat not found")
	}

	db = r.db.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Where("chat_id = ?", chat.ID).Limit(limit).Order("messages.id DESC")
	}).Find(&chat)

	if db.Error != nil {
		return entity.ListIdMessages{}, db.Error
	}

	for i, v := range chat.Messages {
		ListIdMessages[i] = v.ID
	}

	return ListIdMessages, nil
}

func (r *chatRepository) GetMessage(id int) (entity.Message, error) {
	message := entity.Message{}
	db := r.db.Where("id = ?", id).First(&message)
	if db.Error != nil {
		return entity.Message{}, errors.New("message not found")
	}
	return message, nil
}
