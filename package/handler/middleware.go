package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserIndentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		NewErrorResponse(c, 401, "emty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, 401, "No Bearer")
		return
	}

	_, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, 401, "Token not valid")
	}

}
