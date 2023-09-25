package service

import (
	gobackend "github.com/hoach-linux/go-backend"
	"github.com/hoach-linux/go-backend/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{
		repo: repo,
	}
}

func (s *TodoItemService) Create(item gobackend.CrudTodoItem) (int, error) {
	return s.repo.Create(item)
}
func (s *TodoItemService) GetAll(listId int) ([]gobackend.TodoItem, error) {
	return s.repo.GetAll(listId)
}
func (s *TodoItemService) GetById(listId, itemId int) (gobackend.TodoItem, error) {
	return s.repo.GetById(listId, itemId)
}
func (s *TodoItemService) Delete(listId, itemId int) error {
	return s.repo.Delete(listId, itemId)
}
func (s *TodoItemService) Update(listId, itemId int, input gobackend.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(listId, itemId, input)
}
