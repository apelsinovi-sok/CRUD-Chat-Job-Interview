package entity

type Message struct {
	ID            int
	ChatID        int
	MessageAuthor string `json:"message_author" binding:"required"`
	ChatName      string `json:"chat_name" binding:"required"`
	MessageText   string `json:"message_text" binding:"required"`
}
