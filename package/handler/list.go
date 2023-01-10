package handler

import "github.com/gin-gonic/gin"

func (h Handler) createList(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"msg": "hello!",
	})
}

func (h Handler) getAllLists(c *gin.Context) {

}

func (h Handler) getListsById(c *gin.Context) {

}

func (h Handler) updateList(c *gin.Context) {

}

func (h Handler) deleteList(c *gin.Context) {

}
