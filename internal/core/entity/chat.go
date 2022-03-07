package entity

type Chat struct {
	ID         int    `json:"id "`
	ChatName   string `json:"chat_name" binding:"required"`
	ChatAuthor string `json:"chat_author" binding:"required"`
}
