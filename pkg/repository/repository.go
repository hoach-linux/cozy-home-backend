package repository

import (
	"github.com/jmoiron/sqlx"

	gobackend "github.com/hoach-linux/go-backend"
)

type Authorization interface {
	CreateUser(user gobackend.User) (int, error)
	GetUser(username, password string) (gobackend.User, error)
}
type TodoList interface {
	Create(userId int, list gobackend.TodoList) (int, error)
}
type TodoItem interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: newAuthPostgres(db),
	}
}
