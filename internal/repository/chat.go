package repository

import (
	"CRUD-Chat-Test-Task/internal/core/entity"
	"CRUD-Chat-Test-Task/internal/core/interfaces"
	"CRUD-Chat-Test-Task/internal/core/model"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
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
	//Поиск юзера по нику. Если не нашел, создать юзера и создать чат, связать их общим id
	db := r.db.Where("name = ?", chatEntity.ChatAuthor).Limit(1).First(&user)
	if db.RowsAffected == 0 {
		user.ID = userId
		user.Name = chatEntity.ChatAuthor
		r.db.Create(&user)

		//Записать все данные из сущности в модель юзер и чат, для разбиения сущности на таблицы юзер и чат
		chat.ID = uuid.New()
		chat.AuthorID = userId
		chat.ChatName = chatEntity.ChatName
		db = r.db.Create(&chat)
		//В противном случае записать данные из сущности только в модель чат, не создавая нового юзера. Достать id юзера и прикрепить его к чату
	} else {
		chat.ID = uuid.New()
		chat.AuthorID = user.ID
		chat.ChatName = chatEntity.ChatName
		db = r.db.Create(&chat)
	}

	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (r *chatRepository) AddMessage(messageEntity *entity.Message) error {
	currentTime := time.Now()
	chat := model.Chat{}
	user := model.User{}
	message := model.Message{}
	userId := uuid.New()
	//Поиск чата по названию
	db := r.db.Where("chat_name = ?", messageEntity.ChatName).First(&chat)
	if db.Error != nil {
		return errors.New("chat not found")
	}
	message.ChatID = chat.ID

	//Поиск юзера по нику. Если не нашел, создать юзера и создать сообщение, связать их общим id
	db = r.db.Where("name = ?", messageEntity.MessageAuthor).Limit(1).First(&user)
	if db.RowsAffected == 0 {
		user.ID = userId
		user.Name = messageEntity.MessageAuthor
		r.db.Create(&user)
		message.AuthorID = userId
	} else {
		//Если юзер существует, достать его id  связать его с моделью сообщения
		message.AuthorID = user.ID
	}

	message.MessageText = messageEntity.MessageText
	message.CreatedAt = currentTime.Format("2006.01.02 15:04:05")
	message.ID = uuid.New()
	db = r.db.Create(&message)

	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (r *chatRepository) GetListMessages(chatName string, limit int) (entity.ListIdMessages, error) {
	chat := model.Chat{}
	//chatEntity := entity.Chat{}
	ListIdMessages := entity.ListIdMessages{}

	db := r.db.Where("chat_name  = ?", chatName).First(&chat)
	if db.RowsAffected == 0 {
		return entity.ListIdMessages{}, errors.New("chat not found")
	}
	fmt.Println(chat)

	db = r.db.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Where("chat_id = ?", chat.ID).Limit(limit).Order("messages.created_at DESC")
	}).Find(&chat)

	if db.Error != nil {
		return entity.ListIdMessages{}, db.Error
	}

	for i, v := range chat.Messages {
		ListIdMessages[i] = v.ID
	}

	return ListIdMessages, nil
}

func (r *chatRepository) GetMessage(id string) (entity.Message, error) {
	message := model.Message{}
	chat := model.Chat{}
	user := model.User{}
	db := r.db.Where("id = ?", id).First(&message)

	//db = r.db.Preload("Messages", func(db *gorm.DB) *gorm.DB {
	//	return db.Where("id = ?", id)
	//}).Find(&user)

	db = r.db.Where("id = ?", message.AuthorID).First(&user)

	db = r.db.Where("id = ?", message.ChatID).First(&chat)

	messageEntity := entity.Message{
		ID:            id,
		ChatID:        chat.ID,
		MessageAuthor: user.Name,
		ChatName:      chat.ChatName,
		MessageText:   message.MessageText,
		CreatedAt:     message.CreatedAt,
	}

	fmt.Println(messageEntity)
	if db.Error != nil {
		return entity.Message{}, errors.New("message not found")
	}
	return messageEntity, nil
}
