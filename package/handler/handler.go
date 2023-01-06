package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/sign-up", h.signUp)
		authGroup.POST("/sign-in", h.signIn)
	}

	apiGroup := router.Group("/api")
	{
		listGroup := apiGroup.Group("/lists")
		{
			listGroup.POST("/", h.createList)
			listGroup.GET("/", h.getAllLists)
			listGroup.GET("/:id", h.getListsById)
			listGroup.PUT("/:id", h.updateList)
			listGroup.DELETE("/:id", h.deleteList)

			itemsGroup := listGroup.Group("/:id/items")
			{
				itemsGroup.POST("/", h.createItem)
				itemsGroup.GET("/", h.getAllItems)
				itemsGroup.GET("/:itemId", h.getItemsById)
				itemsGroup.PUT("/:itemId", h.updateItem)
				itemsGroup.DELETE("/:itemId", h.deleteItem)
			}
		}
	}

	return router
}
