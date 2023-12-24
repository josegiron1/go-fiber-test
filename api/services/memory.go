package services

import (
	"errors"
	"sync"
	"todo/api/models"
)

type InMemoryTodoService struct {
	todos  map[int]*models.Todo
	mu     sync.Mutex // for concurrent access
	nextID int
}

func NewInMemoryTodoService() *InMemoryTodoService {
	return &InMemoryTodoService{
		todos:  make(map[int]*models.Todo),
		nextID: 1,
	}
}

func (s *InMemoryTodoService) GetTodos() ([]models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todos := make([]models.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, *todo)
	}
	return todos, nil
}

func (s *InMemoryTodoService) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.ID = s.nextID
	s.nextID++

	s.todos[todo.ID] = todo

	return todo, nil
}

func (s *InMemoryTodoService) DeleteTodo(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.todos[id]; !exists {
		return errors.New("todo doesn't exist to be deleted")
	}

	delete(s.todos, id)

	return nil
}
