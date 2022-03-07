package internal

import (
	"CRUD-Chat-Test-Task/internal/core/interfaces"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetListMessages(c *gin.Context) {
	usecase := *c.MustGet("usecase").(*interfaces.Usecase)
	chatName := c.Param("name")
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		c.JSON(BadRequestHttpResponse(errors.New("limit must be a number")))
		return
	}
	ListIdMessages, err := usecase.GetListMessages(chatName, limit)
	if err != nil {
		c.JSON(BadRequestHttpResponse(err))
		return
	}

	c.JSON(SuccessHttpResponse(ListIdMessages))

}
