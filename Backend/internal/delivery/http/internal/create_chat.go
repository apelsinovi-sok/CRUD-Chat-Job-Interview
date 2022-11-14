package internal

import (
	"CRUD-Chat-Test-Task/internal/core/entity"
	"CRUD-Chat-Test-Task/internal/core/interfaces"
	"github.com/gin-gonic/gin"
)

func CreateChat(c *gin.Context) {
	usecase := *c.MustGet("usecase").(*interfaces.Usecase)
	chat := entity.Chat{}
	err := c.ShouldBindJSON(&chat)
	if err != nil {
		c.JSON(BadRequestHttpResponse(err))
		return
	}

	err = usecase.CreateChat(chat)
	if err != nil {
		c.JSON(BadRequestHttpResponse(err))
		return
	}

	c.JSON(SuccessHttpResponse(nil))
}