package gobackend

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}
type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}
type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
type CrudTodoItem struct {
	TodoItem
	ListId string `json:"list_id" data:"list_id" binding:"required"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("Inputs was not change")
	}

	return nil
}
