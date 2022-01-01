package service

import (
	todo "ToDoList"
	"ToDoList/pkg/repository"
)

type ToDoListService struct {
	repo repository.ToDoList
}

func (s *ToDoListService) Update(userId int, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}

func (s *ToDoListService) Delete(userId int, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *ToDoListService) GetById(userId int, listId int) (todo.ToDoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *ToDoListService) GetAll(userId int) ([]todo.ToDoList, error) {
	return s.repo.GetAll(userId)
}

func NewToDoListService(repo repository.ToDoList) *ToDoListService {
	return &ToDoListService{repo: repo}
}

func (s *ToDoListService) Create(userId int, list todo.ToDoList) (int, error) {
	return s.repo.Create(userId, list)
}
