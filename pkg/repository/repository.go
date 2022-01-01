package repository

import (
	todo "ToDoList"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type ToDoList interface {
	Create(id int, list todo.ToDoList) (int, error)
	GetAll(userId int) ([]todo.ToDoList, error)
	GetById(userId int, listId int) (todo.ToDoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input todo.UpdateListInput) error
}

type ToDoItem interface {
	Create(listId int, item todo.ToDoItem) (int, error)
	GetAll(userId int, listId int) ([]todo.ToDoItem, error)
	GetById(userId int, itemId int) (todo.ToDoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ToDoList:      NewToDoListPostgres(db),
		ToDoItem:      NewToDoItemPostgres(db),
	}
}
