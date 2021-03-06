package repository

import (
	"github.com/jmoiron/sqlx"
	todoapi "github.com/klaus-abram/todo-rest-api"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemTable   = "todo_items"
	listsItemsTable = "lists_items"
)

type Authorization interface {
	CreateUser(user todoapi.User) (int, error)
	GetUser(username, password string) (todoapi.User, error)
}

type TodoList interface {
	Create(userId int, list todoapi.TodoList) (int, error)
	GetAll(userId int) ([]todoapi.TodoList, error)
	GetById(userId, listId int) (todoapi.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todoapi.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, input todoapi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todoapi.TodoItem, error)
	GetById(userId, itemId int) (todoapi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todoapi.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
