package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")

		return
	}

	//parse token

	userId, err := h.service.Authorization.ParseToken(headerParts[1])

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "id is not defined")

		return 0, errors.New("user id is not defined")
	}

	idInt, ok := id.(int)

	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "id is not defined")

		return 0, errors.New("user id is not defined")
	}

	return idInt, nil
}
func authCheck(h *Handler, userId, listId int, c *gin.Context) error {
	_, err := h.service.TodoList.GetById(userId, listId)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Permission denied")
	}

	return err
}
