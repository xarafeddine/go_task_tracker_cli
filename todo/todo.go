package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// Constants for todo status
const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

// Todo represents a single task
type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (todo Todo) String() string {
	return fmt.Sprintf("%d: %s <%s> (%s)", todo.ID, todo.Task, todo.Status, todo.UpdatedAt.Format("2006-01-02 15:04:05"))
}

// Todos is a collection of Todo items
type Todos []Todo

// LoadTodos reads the todos from a JSON file
func LoadTodos(filename string) (Todos, error) {
	var todos Todos

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return todos, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read todos from %s: %w", filename, err)
	}

	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("failed to parse todos from %s: %w", filename, err)
	}

	return todos, nil
}

// SaveTodos writes the todos to a JSON file
func (t *Todos) SaveTodos(filename string) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize todos: %w", err)
	}

	return os.WriteFile(filename, data, 0644)
}

// AddTodo adds a new todo item
func (t *Todos) AddTodo(task string) {
	newTodo := Todo{
		ID:        int(time.Now().UnixNano()),
		Task:      task,
		Status:    StatusTodo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	*t = append(*t, newTodo)
}

// UpdateTodo updates the task or status of a todo item
func (t *Todos) UpdateTodo(id int, task, status string) error {
	for i, todo := range *t {
		if todo.ID == id {
			(*t)[i].UpdatedAt = time.Now()
			if task != "" {
				(*t)[i].Task = task
			}
			if status != "" {
				(*t)[i].Status = status
			}
			return nil
		}
	}
	return errors.New("todo not found")
}

// DeleteTodo deletes a todo by ID
func (t *Todos) DeleteTodo(id int) error {
	newTodos := make(Todos, 0)
	for _, todo := range *t {
		if todo.ID != id {
			newTodos = append(newTodos, todo)
		}
	}

	if len(newTodos) == len(*t) {
		return errors.New("todo not found")
	}

	*t = newTodos
	return nil
}

// ListTodos lists todos by status
func (t *Todos) ListTodos(status string) {
	if len(*t) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, todo := range *t {
		if status == "" || todo.Status == status {
			fmt.Println(todo)
		}
	}
}
