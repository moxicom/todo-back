package service

import (
	"github.com/moxicom/todo-back/models"
	"github.com/moxicom/todo-back/pkg/repository"
)

type itemService struct {
	repository repository.Item
}

func newItemService(repository repository.Item) *itemService {
	return &itemService{
		repository: repository,
	}
}

func (s *itemService) CheckList(userId, listId int) error {
	return s.repository.CheckList(userId, listId)
}

func (s *itemService) Create(userId, listId int, item models.Item) (int, error) {
	if err := s.CheckList(userId, listId); err != nil {
		return 0, err
	}
	return s.repository.Create(listId, item)
}

func (s *itemService) GetAll(userId, listId int) ([]models.Item, error) {
	if err := s.CheckList(userId, listId); err != nil {
		return []models.Item{}, err
	}
	return s.repository.GetAll(listId)
}

func (s *itemService) GetById(userId, listId, itemId int) (models.Item, error) {
	if err := s.CheckList(userId, listId); err != nil {
		return models.Item{}, err
	}
	return s.repository.GetById(listId, itemId)
}

func (s *itemService) Update(userId, listId, itemId int, input models.Item) error {
	if err := s.CheckList(userId, listId); err != nil {
		return err
	}
	return s.repository.Update(listId, itemId, input)
}

func (s *itemService) Delete(userId, listId, itemId int) error {
	if err := s.CheckList(userId, listId); err != nil {
		return err
	}
	return s.repository.Delete(listId, itemId)
}
