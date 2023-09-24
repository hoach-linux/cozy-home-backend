package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gobackend "github.com/hoach-linux/go-backend"
)

type getAllItemsResponse struct {
	Data []gobackend.TodoItem `json:"data"`
}

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	var input gobackend.CrudTodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	inputListId, err := strconv.Atoi(input.ListId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.service.TodoList.GetById(userId, inputListId)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Permission denied")
		return
	}

	id, err := h.service.TodoItem.Create(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getItems(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Query("list_id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "query param is not valid")
		return
	}

	_, err = h.service.TodoList.GetById(userId, listId)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Permission denied")
		return
	}

	items, err := h.service.TodoItem.GetAll(listId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id is not valid")
		return
	}

	listId, err := strconv.Atoi(c.Query("list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "query param is not valid")
		return
	}

	_, err = h.service.TodoList.GetById(userId, listId)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Permission denied")
		return
	}

	item, err := h.service.TodoItem.GetById(listId, itemId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}
func (h *Handler) updateItem(c *gin.Context) {

}
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id is not valid")
		return
	}

	listId, err := strconv.Atoi(c.Query("list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "query param is not valid")
		return
	}

	_, err = h.service.TodoList.GetById(userId, listId)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Permission denied")
		return
	}

	err = h.service.TodoItem.Delete(listId, itemId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
