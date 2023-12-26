package service

import (
	"github.com/moxicom/todo-back/models"
	"github.com/moxicom/todo-back/pkg/repository"
)

type TodoService struct {
	repository repository.TodoList
}

func NewTodoService(repository repository.TodoList) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

func (s *TodoService) Create(userId int, list models.TodoList) (int, error) {
	return s.repository.Create(userId, list)
}

func (s *TodoService) GetAll(userId int) ([]models.TodoList, error) {
	return s.repository.GetAll(userId)
}
