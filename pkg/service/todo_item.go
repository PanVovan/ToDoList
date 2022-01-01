package service

import (
	todo "ToDoList"
	"ToDoList/pkg/repository"
)

type ToDoItemService struct {
	repo     repository.ToDoItem
	listRepo repository.ToDoList
}

func (s *ToDoItemService) Update(userId, itemId int, input todo.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, itemId, input)
}

func (s *ToDoItemService) Delete(userId int, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *ToDoItemService) GetById(userId int, itemId int) (todo.ToDoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *ToDoItemService) GetAll(userId int, listId int) ([]todo.ToDoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func NewToDoItemService(repo repository.ToDoItem, listRepo repository.ToDoList) *ToDoItemService {
	return &ToDoItemService{repo: repo, listRepo: listRepo}
}

func (s *ToDoItemService) Create(userId, listId int, item todo.ToDoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, item)
}
