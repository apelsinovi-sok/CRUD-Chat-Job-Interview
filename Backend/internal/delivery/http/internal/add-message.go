package internal

import (
	"CRUD-Chat-Test-Task/internal/core/entity"
	"CRUD-Chat-Test-Task/internal/core/interfaces"
	"github.com/gin-gonic/gin"
)

func AddMessage(c *gin.Context) {
	usecase := *c.MustGet("usecase").(*interfaces.Usecase)
	message := entity.Message{}
	err := c.ShouldBindJSON(&message)
	if err != nil {
		c.JSON(BadRequestHttpResponse(err))
		return
	}

	err = usecase.AddMessage(message)
	if err != nil {
		c.JSON(BadRequestHttpResponse(err))
		return
	}

	c.JSON(SuccessHttpResponse(nil))
}
