package internal

import (
	"CRUD-Chat-Test-Task/internal/core/interfaces"
	"github.com/gin-gonic/gin"
)

func GetMessage(c *gin.Context) {
	usecase := *c.MustGet("usecase").(*interfaces.Usecase)
	messageId := c.Param("id")
	message, err := usecase.GetMessage(messageId)
	if err != nil {
		c.JSON(BadRequestHttpResponse(err))
		return
	}

	c.JSON(SuccessHttpResponse(message))

}
