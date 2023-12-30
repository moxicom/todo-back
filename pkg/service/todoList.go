package service

import (
	"github.com/moxicom/todo-back/models"
	"github.com/moxicom/todo-back/pkg/repository"
)

type todoService struct {
	repository repository.TodoList
}

func newTodoService(repository repository.TodoList) *todoService {
	return &todoService{
		repository: repository,
	}
}

func (s *todoService) Create(userId int, list models.TodoList) (int, error) {
	return s.repository.Create(userId, list)
}

func (s *todoService) GetAll(userId int) ([]models.TodoList, error) {
	return s.repository.GetAll(userId)
}

func (s *todoService) GetById(userId, listId int) (models.TodoList, error) {
	return s.repository.GetById(userId, listId)
}

func (s *todoService) Update(userId, listId int, input models.TodoList) error {
	return s.repository.Update(userId, listId, input)
}

func (s *todoService) Delete(userId, listId int) error {
	return s.repository.Delete(userId, listId)
}
