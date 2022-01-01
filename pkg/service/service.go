package service

import (
	todo "ToDoList"
	"ToDoList/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	Create(userId int, list todo.ToDoList) (int, error)
	GetAll(userId int) ([]todo.ToDoList, error)
	GetById(userId int, listId int) (todo.ToDoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input todo.UpdateListInput) error
}

type ToDoItem interface {
	Create(userId int, listId int, input todo.ToDoItem) (int, error)
	GetAll(userId int, listId int) ([]todo.ToDoItem, error)
	GetById(userId int, itemId int) (todo.ToDoItem, error)
	Delete(userId int, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ToDoList:      NewToDoListService(repos.ToDoList),
		ToDoItem:      NewToDoItemService(repos.ToDoItem, repos.ToDoList),
	}
}
