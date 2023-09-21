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
	}
}
