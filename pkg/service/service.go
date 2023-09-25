package service

import (
	gobackend "github.com/hoach-linux/go-backend"
	"github.com/hoach-linux/go-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user gobackend.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list gobackend.TodoList) (int, error)
	GetAll(userId int) ([]gobackend.TodoList, error)
	GetById(userId, listId int) (gobackend.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input gobackend.UpdateListInput) error
}
type TodoItem interface {
	Create(item gobackend.CrudTodoItem) (int, error)
	GetAll(listId int) ([]gobackend.TodoItem, error)
	GetById(listId, itemId int) (gobackend.TodoItem, error)
	Delete(listId, itemId int) error
	Update(listId, itemId int, input gobackend.UpdateItemInput) error
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewItemService(repos.TodoItem),
	}
}
