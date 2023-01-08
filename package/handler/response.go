package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode uint16, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(int(statusCode), Error{Message: message})
}
