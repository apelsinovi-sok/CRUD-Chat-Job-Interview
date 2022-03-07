package entity

type Chat struct {
	ID     int    `json:"id"`
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
}
