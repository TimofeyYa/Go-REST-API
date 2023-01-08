package handler

import (
	todo "todo/study"

	"github.com/gin-gonic/gin"
)

func (h Handler) signUp(c *gin.Context) {
	var InputData todo.RegUser
	if err := c.BindJSON(&InputData); err != nil {
		NewErrorResponse(c, 400, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(InputData)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}

func (h Handler) signIn(c *gin.Context) {
	var InputData todo.User
	if err := c.BindJSON(&InputData); err != nil {
		NewErrorResponse(c, 400, err.Error())
		return
	}

	token, err := h.service.Authorization.LoginUser(InputData)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, map[string]interface{}{
		"token": token,
	})
}
