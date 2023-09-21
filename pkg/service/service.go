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
}
type TodoItem interface {
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		TodoList: NewTodoListService(repos.TodoList),
	}
}
