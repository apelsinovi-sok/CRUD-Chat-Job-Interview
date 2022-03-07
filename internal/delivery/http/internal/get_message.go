package internal

import (
	"CRUD-Chat-Test-Task/internal/core/interfaces"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetMessage(c *gin.Context) {
	usecase := *c.MustGet("usecase").(*interfaces.Usecase)
	messageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(BadRequestHttpResponse(errors.New("invalid argument")))
		return
	}
	message, err := usecase.GetMessage(messageId)
	if err != nil {
		c.JSON(BadRequestHttpResponse(err))
		return
	}

	c.JSON(SuccessHttpResponse(message))

}
