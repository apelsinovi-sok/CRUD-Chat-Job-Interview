package internal

import (
	"CRUD-Chat-Test-Task/internal/core/entity"
	"net/http"
)

func SuccessHttpResponse(data interface{}) (int, entity.Dict) {
	if data == nil {
		return http.StatusOK,
			entity.Dict{"success": true}
	}

	return http.StatusOK,
		entity.Dict{
			"success": true,
			"data":    data,
		}
}

func BadRequestHttpResponse(err error) (int, entity.Dict) {

	return http.StatusBadRequest,
		entity.Dict{
			"success": false,
			"data":    err.Error(),
		}
}
