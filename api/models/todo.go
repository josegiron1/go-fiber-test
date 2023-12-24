package models

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type TodoService interface {
	CreateTodo() (*Todo, error)

	GetTodos() (*[]Todo, error)

	DeleteTodo() error
}
