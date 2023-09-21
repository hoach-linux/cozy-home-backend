package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gobackend "github.com/hoach-linux/go-backend"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "id is not defined")

		return
	}

	var input gobackend.TodoList

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}
	

}
func (h *Handler) getLists(c *gin.Context) {
	id, _ := c.Get(userCtx)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getListById(c *gin.Context) {

}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
